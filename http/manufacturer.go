package http

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kmcclive/goapipattern"
)

type ManufacturerController struct {
	service goapipattern.ManufacturerService
}

func NewManufacturerController(service goapipattern.ManufacturerService) *ManufacturerController {
	return &ManufacturerController{
		service: service,
	}
}

func (c *ManufacturerController) Create(ctx *gin.Context) {
	var manufacturer goapipattern.Manufacturer

	if err := ctx.ShouldBind(&manufacturer); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	if err := c.service.Create(&manufacturer); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, manufacturer)
}

func (c *ManufacturerController) Delete(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	if err := c.service.Delete(uint(id)); err != nil {
		if errors.Is(err, goapipattern.ErrNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, NewError(
				fmt.Sprintf("A manufacturer with id %v was not found", id),
			))
		} else {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *ManufacturerController) FetchByID(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	manufacturer, err := c.service.FetchByID(uint(id))
	if err != nil {
		if errors.Is(err, goapipattern.ErrNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, NewError(
				fmt.Sprintf("A manufacturer with id %v was not found", id),
			))
		} else {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, manufacturer)
}

func (c *ManufacturerController) List(ctx *gin.Context) {
	manufacturers, err := c.service.List()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, manufacturers)
}

func (c *ManufacturerController) Update(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	var manufacturer goapipattern.Manufacturer

	if err := ctx.ShouldBind(&manufacturer); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	manufacturer.ID = uint(id)

	if err := c.service.Update(&manufacturer); err != nil {
		if errors.Is(err, goapipattern.ErrNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, NewError(
				fmt.Sprintf("A manufacturer with id %v was not found", id),
			))
		} else {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, manufacturer)
}
