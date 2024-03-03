package rest

import (
	driver "go-blog/Driver"
	"go-blog/gateway"
	"go-blog/usecase"
	"net/http"

	"github.com/labstack/echo"
)

func GetBlogsRest(e *echo.Echo) {
	e.GET("/blogs", getBlogsHandler)
}

func getBlogsHandler(c echo.Context) error {
	driver := driver.NewBlogsDriver()
	gateway := gateway.NewGetBlogsGateway(driver)
	usecase := usecase.NewGetBlogsUsecase(gateway)
	blogs, _ := usecase.Execute()
	return c.JSON(http.StatusOK, blogs)
}
