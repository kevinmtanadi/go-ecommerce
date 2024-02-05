package controllers

import (
	"go-backend-template/src/constants"
	"go-backend-template/src/dtos"
	"go-backend-template/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type UserController interface {
	Register(c echo.Context) error
	Verify(c echo.Context) error
}

type UserControllerImpl struct {
	service *services.Service
}

func NewUserController(ioc di.Container) UserController {
	return &UserControllerImpl{
		service: ioc.Get(constants.SERVICE).(*services.Service),
	}
}

// @Summary Register creates a new user
// @Description Register creates a new user
// @Tags user
// @Router /user/register [post]
// @Accept json
// @Produce json
// @Param Content-Type header string true "content type request" Enums(application/json)
// @Param qs body dtos.RegisterReq true "Register creates a new user"
func (h *UserControllerImpl) Register(c echo.Context) error {
	var (
		params   *dtos.RegisterReq = new(dtos.RegisterReq)
		response dtos.Response
		data     dtos.RegisterRes
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

	data, err = h.service.User.Register(ctx, *params)
	if err != nil {
		response = dtos.NewResponse(nil, err, err.Error())
		return c.JSON(response.Code, response)
	}

	return c.JSON(http.StatusOK, dtos.NewResponse(data, err, "Success"))
}

// @Summary Verify a user
// @Description Verify a user
// @Tags user
// @Router /user/verify [put]
// @Accept json
// @Produce json
// @Param Content-Type header string true "content type request" Enums(application/json)
// @Param qs query dtos.VerifyUserReq true "Verify a user"
func (h *UserControllerImpl) Verify(c echo.Context) error {
	var (
		params   *dtos.VerifyUserReq = new(dtos.VerifyUserReq)
		response dtos.Response
		data     dtos.VerifyUserRes
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

	data, err = h.service.User.Verify(ctx, *params)
	if err != nil {
		response = dtos.NewResponse(nil, err, err.Error())
		return c.JSON(response.Code, response)
	}

	return c.JSON(http.StatusOK, dtos.NewResponse(data, err, "Success"))
}
