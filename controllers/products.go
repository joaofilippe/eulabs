package controllers

import (
	"fmt"
	"net/http"

	"github.com/joaofilippe/go-products/business"
	"github.com/joaofilippe/go-products/database"
	"github.com/joaofilippe/go-products/models"
	"github.com/labstack/echo/v4"
)

type ErrorMessage struct {
	Err     error  `json:"err"`
	Message string `json:"message"`
}

func Create(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		errorMessage := &ErrorMessage{
			Err:     err,
			Message: "Algo deu errado criação do produto",
		}
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, errorMessage)
	}

	messageBusiness, errorBusinnes := business.Create(&product)

	if errorBusinnes != nil {
		return c.JSON(errorBusinnes.StatusCode, errorBusinnes.Message)
	}

	return c.JSON(http.StatusCreated, messageBusiness)
}

func GetById(c echo.Context) error {
	var product models.Product
	id := c.Param("id")
	database.DB.Where("id = ?", id).First(&product)

	if product.ID == "" {
		return c.JSON(http.StatusNotFound,
			"Produto não localizado no banco de dados. Verifique o id informado.")
	}

	return c.JSON(http.StatusFound, &product)

}

func Update(c echo.Context) error {
	var product models.Product
	id := c.Param("id")

	database.DB.Where("id = ?", id).First(&product)

	if product.ID == "" {
		return c.JSON(http.StatusNotFound,
			"Produto não localizado no banco de dados. Verifique o id informado.")
	}

	if err := c.Bind(&product); err != nil {
		errorMessage := &ErrorMessage{
			Err:     err,
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, errorMessage)
	}

	database.DB.Model(&product).UpdateColumns(product)

	return c.JSON(http.StatusAccepted, 
			      "Informações do produto alteradas com sucesso!")
}

func Delete(c echo.Context) error {
	var product models.Product
	id := c.Param("id")

	database.DB.Where("id = ?", id).Delete(&product)

	return c.JSON(http.StatusOK, 
		          "Produto deletado com sucesso")
}
