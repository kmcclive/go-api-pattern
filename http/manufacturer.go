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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.Create(&manufacturer); err != nil {
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, manufacturer)
}

func (c *ManufacturerController) Delete(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.service.Delete(uint(id))

	ctx.Status(http.StatusNoContent)
}

func (c *ManufacturerController) FetchByID(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	manufacturer, err := c.service.FetchByID(uint(id))
	if err != nil {
		ctx.Error(err)
		return
	}
	if manufacturer == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("A manufacturer with id %v was not found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, manufacturer)
}

func (c *ManufacturerController) List(ctx *gin.Context) {
	manufacturers, err := c.service.List()
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, manufacturers)
}

func (c *ManufacturerController) Update(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var manufacturer goapipattern.Manufacturer

	if err := ctx.ShouldBind(&manufacturer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	manufacturer.ID = uint(id)

	if err := c.service.Update(&manufacturer); err != nil {
		if errors.Is(err, goapipattern.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("A manufacturer with id %v was not found", id),
			})
		} else {
			ctx.Error(err)
		}
		return
	}

	ctx.JSON(http.StatusOK, manufacturer)
}
