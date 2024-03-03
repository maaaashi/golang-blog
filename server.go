package main

import (
	"go-blog/rest"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	rest.GetBlogsRest(e)
	e.Logger.Fatal(e.Start(":1323"))
}
