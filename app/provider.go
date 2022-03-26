package app

import (
	"github.com/labstack/echo/v4"
	"nes/app/api"
	"nes/app/dao"
	"nes/config"
)

func Init(ec *echo.Echo, conf *config.Conf) {
	db := dao.NewDbClient(&conf.Database)

	ep := &api.EndPoint{
		DB: db,
	}
	g := ec.Group("/api")
	ep.InitCustomerRouter(g)
}
