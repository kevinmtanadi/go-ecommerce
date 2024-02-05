package main

import (
	"go-backend-template/src/controllers"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type Route struct {
	app        *echo.Echo
	controller *controllers.Controller
	router     *echo.Group
}

func NewRoute(app *echo.Echo, ioc di.Container) *Route {
	return &Route{
		app:        app,
		controller: controllers.NewController(ioc),
		router:     app.Group(""),
	}
}

func (r *Route) Init() {
	r.Template()
	r.MySQL()
	r.User()
}

func (r *Route) Template() {
	r.router.GET("/template", r.controller.Template.TemplateFunc)
}

func (r *Route) MySQL() {
	mySql := r.router.Group("/mysql")
	mySql.GET("", r.controller.MySQL.MySQLQueryEditor)
	mySql.POST("/query", r.controller.MySQL.RunQuery)
	mySql.GET("/table", r.controller.MySQL.GetTableNames)
}

func (r *Route) User() {
	user := r.router.Group("/user")
	user.POST("/register", r.controller.User.Register)
	user.PUT("/verify", r.controller.User.Verify)
}
