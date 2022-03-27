package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"nes/app"
	"nes/assets"
	"nes/config"

	mw "github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("网络评估系统")

	conf := &config.Conf{}
	config.GetConfig(conf)

	e := echo.New()

	e.Use(mw.CORSWithConfig(mw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAuthorization},
	}))

	app.Init(e, conf)

	e.StaticFS("/", assets.FileSystem)

	e.Logger.Fatal(e.Start(":8080"))
}
