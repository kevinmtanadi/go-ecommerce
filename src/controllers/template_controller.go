package controllers

import (
	"go-backend-template/src/constants"
	"go-backend-template/src/dtos"
	"go-backend-template/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type TemplateController interface {
	TemplateFunc(c echo.Context) error
}

type TemplateControllerImpl struct {
	service *services.Service
}

func NewTemplateController(ioc di.Container) TemplateController {
	return &TemplateControllerImpl{
		service: ioc.Get(constants.SERVICE).(*services.Service),
	}
}

// @Summary Template function
// @Description Template function
// @Tags template
// @Router /template [get]
// @Accept json
// @Produce json
// @Param Content-Type header string true "content type request" Enums(application/json)
// @Param qs query dtos.TemplateDtoReq true "Template function"
func (h *TemplateControllerImpl) TemplateFunc(c echo.Context) error {
	var (
		params   *dtos.TemplateDtoReq = new(dtos.TemplateDtoReq)
		response dtos.Response
		data     dtos.TemplateDtoRes
		err      error
	)

	err = (&echo.DefaultBinder{}).BindQueryParams(c, params)
	if err != nil {
		response = dtos.NewResponse(nil, err, "Error while binding request")
		return c.JSON(response.Code, response)
	}

	if err = c.Bind(&params); err != nil {
		response = dtos.NewResponse(nil, err, "Error while binding request")
		return c.JSON(200, response)
	}

	ctx := c.Request().Context()

	data, err = h.service.Template.TemplateFunc(ctx, *params)
	if err != nil {
		response = dtos.NewResponse(nil, err, err.Error())
		return c.JSON(response.Code, response)
	}

	return c.JSON(http.StatusOK, dtos.NewResponse(data, err, "Success"))
}
