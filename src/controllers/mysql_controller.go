package controllers

import (
	"go-backend-template/src/constants"
	"go-backend-template/src/dtos"
	"go-backend-template/src/services"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type MySQLController interface {
	MySQLQueryEditor(c echo.Context) error
	RunQuery(c echo.Context) error
	GetTableNames(c echo.Context) error
}

type MySQLControllerImpl struct {
	service *services.Service
}

func NewMySQLController(ioc di.Container) MySQLController {
	return &MySQLControllerImpl{
		service: ioc.Get(constants.SERVICE).(*services.Service),
	}
}

func (h *MySQLControllerImpl) MySQLQueryEditor(c echo.Context) error {
	tmpl, err := template.ParseFiles("./src/template/mysql_query_editor.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response().Writer, nil)
}

func (h *MySQLControllerImpl) RunQuery(c echo.Context) error {
	type Query struct {
		Body string `json:"body"`
	}

	query := Query{}
	err := c.Bind(&query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.NewResponse(nil, err, err.Error()))
	}

	result, err := h.service.MySQL.RunQuery(query.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.NewResponse(nil, err, err.Error()))
	}

	templateData := map[string]interface{}{
		"data":    result,
		"message": "",
	}

	return c.JSON(http.StatusOK, templateData)
}

func (h *MySQLControllerImpl) GetTableNames(c echo.Context) error {
	query := `SHOW TABLES`

	result, err := h.service.MySQL.RunQuery(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.NewResponse(nil, err, err.Error()))
	}

	return c.JSON(http.StatusOK, result)
}
