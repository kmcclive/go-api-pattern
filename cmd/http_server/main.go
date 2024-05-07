package main

import (
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/kmcclive/goapipattern"
	"github.com/kmcclive/goapipattern/http"
	"github.com/kmcclive/goapipattern/sql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("catalog.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&goapipattern.Manufacturer{}, &goapipattern.Product{})

	r := gin.Default()

	manufacturerService := sql.NewManufacturerService(db)
	manufacturerController := http.NewManufacturerController(manufacturerService)
	manufacturerGroup := r.Group("/manufacturers")
	{
		manufacturerGroup.GET("", manufacturerController.List)
		manufacturerGroup.POST("", manufacturerController.Create)
		manufacturerGroup.DELETE("/:id", manufacturerController.Delete)
		manufacturerGroup.GET("/:id", manufacturerController.FetchByID)
		manufacturerGroup.PUT("/:id", manufacturerController.Update)
	}

	r.Run()
}
