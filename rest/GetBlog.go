package rest

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetBlogsRest(e *echo.Echo) {
	e.GET("/blogs", getBlogsHandler)
}

func getBlogsHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
