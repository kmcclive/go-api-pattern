package http

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kmcclive/goapipattern"
)

type ProductController struct {
	service goapipattern.ProductService
}

func NewProductController(service goapipattern.ProductService) *ProductController {
	return &ProductController{
		service: service,
	}
}

func (c *ProductController) Create(ctx *gin.Context) {
	var product goapipattern.Product

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	if err := c.service.Create(&product); err != nil {
		if errors.Is(err, goapipattern.ErrParentNotFound) {
			ctx.AbortWithStatusJSON(http.StatusConflict, NewError(
				fmt.Sprintf("A manufacturer with id %v was not found", product.ManufacturerID),
			))
		} else {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) Delete(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	if err := c.service.Delete(uint(id)); err != nil {
		if errors.Is(err, goapipattern.ErrNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, NewError(
				fmt.Sprintf("A product with id %v was not found", id),
			))
		} else {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *ProductController) FetchByID(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	product, err := c.service.FetchByID(uint(id))
	if err != nil {
		if errors.Is(err, goapipattern.ErrNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, NewError(
				fmt.Sprintf("A product with id %v was not found", id),
			))
		} else {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) List(ctx *gin.Context) {
	var products *[]goapipattern.Product
	var err error

	if idString := ctx.Query("manufacturerId"); idString != "" {
		var id uint64

		id, err = strconv.ParseUint(idString, 10, 32)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, NewError(err.Error()))
			return
		}

		products, err = c.service.SearchByManufacturer(uint(id))
	} else {
		products, err = c.service.List()
	}

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (c *ProductController) Update(ctx *gin.Context) {
	idString := ctx.Param("id")

	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	var product goapipattern.Product

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	product.ID = uint(id)

	if err := c.service.Update(&product); err != nil {
		if errors.Is(err, goapipattern.ErrNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, NewError(
				fmt.Sprintf("A product with id %v was not found", id),
			))
		} else if errors.Is(err, goapipattern.ErrParentNotFound) {
			ctx.AbortWithStatusJSON(http.StatusConflict, NewError(
				fmt.Sprintf("A manufacturer with id %v was not found", product.ManufacturerID),
			))
		} else {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, product)
}
