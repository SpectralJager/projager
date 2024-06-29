package main

import (
	"context"
	"log"
	"os"

	vindex "github.com/SpectralJager/projager/view/index"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	generalHandler := GeneralHandler{}

	app.StaticFS("/static", os.DirFS("./public"))

	app.GET("/", generalHandler.Index)

	log.Fatal(app.Start(":8080"))
}

type GeneralHandler struct{}

func (h GeneralHandler) Index(ctx echo.Context) error {
	return vindex.Page().Render(context.TODO(), ctx.Response())
}
