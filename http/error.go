package http

import (
	"github.com/gin-gonic/gin"
)

func NewError(error string) gin.H {
	return gin.H{
		"error": error,
	}
}
