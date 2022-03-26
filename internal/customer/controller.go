package customer

import (
	"github.com/labstack/echo/v4"
	"nes/pkg/apidocs"
)

type Customer struct {
	Company string `json:"company"`
	Name    string `json:"name"`
	Mobile  string `json:"mobile"`
}

var customers = make([]Customer, 0)

func GetCustomer(c echo.Context) error {
	return c.JSON(apidocs.Success(customers))
}

func CreateCustomer(c echo.Context) error {
	customer := new(Customer)
	if err := c.Bind(customer); err != nil {
		return c.JSON(apidocs.Err(err.Error()))
	}

	customers = append(customers, *customer)

	return c.JSON(apidocs.Success(*customer))

}
