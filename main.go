package main

import (
	"github.com/joaofilippe/go-products/database"
	"github.com/joaofilippe/go-products/migrations"
	"github.com/joaofilippe/go-products/routes"
)

func main() {
	database.Connect()
	migrations.CreateProducstTable()
	routes.HandleRequests()
}
