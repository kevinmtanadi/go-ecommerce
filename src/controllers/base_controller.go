package controllers

import (
	"github.com/sarulabs/di"
)

type Controller struct {
	Template TemplateController
	MySQL    MySQLController
	User     UserController
}

func NewController(ioc di.Container) *Controller {
	return &Controller{
		Template: NewTemplateController(ioc),
		MySQL:    NewMySQLController(ioc),
		User:     NewUserController(ioc),
	}
}
