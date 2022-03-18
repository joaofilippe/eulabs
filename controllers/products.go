package controllers

import (
	"fmt"
	"net/http"

	"github.com/joaofilippe/go-products/database"
	"github.com/joaofilippe/go-products/models"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		errorMessage := &ErrorMessage{
			err:     err,
			message: "Algo deu errado criação do usuário",
		}
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, errorMessage)
	}

	database.DB.Create(&product)

	return c.JSON(http.StatusCreated, product)
}

func GetById(c echo.Context) error {
	var product models.Product
	id := c.Param("id")
	database.DB.First(&product, id)

	return c.JSON(http.StatusFound, product)

}

func Update(c echo.Context) error {
	var product models.Product
	id := c.Param("id")

	database.DB.First(&product, id)

	if err := c.Bind(&product); err != nil {
		errorMessage := &ErrorMessage{
			err:     err,
			message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, errorMessage)
	}

	database.DB.Model(&product).UpdateColumns(product)

	return c.JSON(http.StatusAccepted, "Informações do usuário alteradas com sucesso!")
}

func Delete(c echo.Context) error {
	productId := c.Param("id")

	database.DB.Delete(&models.Product{}, productId)

	return c.JSON(http.StatusOK, "Produto deletado com sucesso")
}

type ErrorMessage struct {
	err     error
	message string
}
