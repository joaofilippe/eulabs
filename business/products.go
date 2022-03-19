package business

import (
	"net/http"

	"github.com/joaofilippe/go-products/database"
	"github.com/joaofilippe/go-products/models"
	"github.com/lithammer/shortuuid"
)

type SucessMessage struct {
	StatusCode int    `json: "statusCode"`
	Message    string `json:"message"`
	Data       string `json:"data"`
}

type ErrorMessage struct {
	StatusCode int    `json: "statusCode"`
	Message    string `json: "message"`
}

func Create(product *models.Product) (*SucessMessage, *ErrorMessage) {
	id := shortuuid.New()

	product.ID = id

	if product.Name == "" {
		m := &ErrorMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Você deve atribuir um nome ao produto",
		}

		return nil, m
	}

	if product.Brand == "" {
		m := &ErrorMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Você deve indicar o nome da marca do produto.",
		}

		return nil, m
	}

	if product.Price == 0 {
		m := &ErrorMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Você precisa inserir o preço do produto.",
		}

		return nil, m
	}

	if product.Quantity == 0 {
		m := &ErrorMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "Você precisa indicar a quantidade do produto no estoque.",
		}

		return nil, m
	}

	database.DB.Create(&product)

	m := &SucessMessage{
		StatusCode: http.StatusCreated,
		Message:    "Produto criado com sucesso!!",
		Data:       id,
	}

	return m, nil
}
