package migrations

import (
	"github.com/joaofilippe/go-products/database"
	"github.com/joaofilippe/go-products/models"
)

func CreateProducstTable() {
	database.DB.AutoMigrate(&models.Product{})
}