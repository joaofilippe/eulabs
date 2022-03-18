package routes

import (
	"net/http"

	"github.com/joaofilippe/go-products/controllers"
	"github.com/labstack/echo/v4"
)

func HandleRequests() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/", controllers.Create)
	e.GET("/:id", controllers.GetById)
	e.PUT("/:id", controllers.Update)
	e.DELETE("/:id", controllers.Delete)
	e.Logger.Fatal(e.Start(":1323"))
}