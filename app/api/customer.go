package api

import (
	"github.com/labstack/echo/v4"
	"nes/app/model"
	"nes/pkg/apidocs"
)

func (ep *EndPoint) InitCustomerRouter(g *echo.Group) {
	sg := g.Group("/customers")
	sg.GET("", ep.ListCustomers)
	sg.POST("", ep.CreateCustomer)
}

func (ep *EndPoint) ListCustomers(c echo.Context) error {
	var customers []model.Customer
	if err := ep.DB.List(&customers); err != nil {
		return c.JSON(apidocs.Err(err.Error()))
	}
	return c.JSON(apidocs.Success(customers))
}

func (ep *EndPoint) CreateCustomer(c echo.Context) error {
	customer := new(model.Customer)
	if err := c.Bind(customer); err != nil {
		return c.JSON(apidocs.Err(err.Error()))
	}

	ep.DB.Create(customer)

	return c.JSON(apidocs.Success(*customer))
}
