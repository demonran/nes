package app

import (
	"github.com/labstack/echo/v4"
	"nes/app/api"
	"nes/app/dao"
	"nes/config"
	"nes/pkg/apidocs"
)

func Init(ec *echo.Echo, conf *config.Conf) {
	db := dao.NewDbClient(&conf.Database)

	ep := &api.EndPoint{
		DB: db,
	}

	g := ec.Group("/api")
	g.GET("/ping", func(c echo.Context) error {
		return c.JSON(apidocs.Success(map[string]string{"status": "ok"}))
	})
	ep.InitCustomerRouter(g)
	ep.InitEvaluationRouter(g)
}
