package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"nes/assets"
	"nes/config"
	"nes/internal/customer"
)

func main() {
	fmt.Println("网络评估系统")

	conf := &config.Conf{}
	config.GetConfig(conf)

	e := echo.New()

	e.POST("/customers", customer.CreateCustomer)
	e.GET("/customers", customer.GetCustomer)

	e.StaticFS("/", assets.FileSystem)

	e.Logger.Fatal(e.Start(":8080"))
}
