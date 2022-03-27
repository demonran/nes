package api

import (
	"github.com/labstack/echo/v4"
	"nes/app/model"
	"nes/pkg/apidocs"
	"nes/pkg/snowflake"
)

func (ep *EndPoint) InitEvaluationRouter(g *echo.Group) {
	sg := g.Group("/evaluation")
	sg.POST("/submit", ep.SubmitEvaluationForm)
}

func (ep *EndPoint) SubmitEvaluationForm(c echo.Context) error {
	var evaluation = new(model.Evaluation)

	if err := c.Bind(evaluation); err != nil {
		return c.JSON(apidocs.Err(err.Error()))
	}

	evaluation.SetId(snowflake.NextId())

	if evaluation.GeneralInfo != nil {
		evaluation.GeneralInfo.SetId(snowflake.NextId())
	}

	for _, item := range evaluation.SiteInfos {
		item.SetId(snowflake.NextId())
	}
	ep.DB.Save(evaluation)
	return c.JSON(apidocs.Success(evaluation))
}
