package main

import (
	"github.com/glebarez/sqlite"
	"github.com/kmcclive/goapipattern"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("catalog.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&goapipattern.Manufacturer{})
}
