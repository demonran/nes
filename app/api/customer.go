package api

import (
	"github.com/labstack/echo/v4"
	"nes/app/model"
	"nes/pkg/apidocs"
	"nes/pkg/snowflake"
)

func (ep *EndPoint) InitCustomerRouter(g *echo.Group) {
	sg := g.Group("/customers")
	sg.GET("", ep.ListCustomers)
	sg.POST("", ep.CreateCustomer)
}

func (ep *EndPoint) ListCustomers(c echo.Context) error {
	var models []model.Customer
	err := ep.DB.FindAll(&models)
	if err != nil {
		return c.JSON(apidocs.Err(err.Error()))
	}
	return c.JSON(apidocs.Success(models))
}

func (ep *EndPoint) CreateCustomer(c echo.Context) error {
	customer := new(model.Customer)
	if err := c.Bind(customer); err != nil {
		return c.JSON(apidocs.Err(err.Error()))
	}

	customer.SetId(snowflake.NextId())
	ep.DB.Save(customer)

	return c.JSON(apidocs.Success(*customer))
}
