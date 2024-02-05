package main

import (
	"go-backend-template/src/constants"
	"go-backend-template/src/controllers"
	"go-backend-template/src/libraries"
	"go-backend-template/src/pkg"
	"go-backend-template/src/services"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type Module struct{}

func (m *Module) New(app *echo.Echo) {
	s := NewService(app)
	s.UseMiddlewares()

	ioc := m.NewIOC()

	r := NewRoute(app, ioc)
	r.Init()
}

func (m *Module) NewIOC() di.Container {
	builder, _ := di.NewBuilder()
	_ = builder.Add(
		di.Def{
			Name: constants.MYSQL,
			Build: func(ctn di.Container) (interface{}, error) {
				db, err := pkg.NewMysqlClient()
				return db, err
			},
		},
		di.Def{
			Name: constants.CONTROLLER,
			Build: func(ctn di.Container) (interface{}, error) {
				return controllers.NewController(builder.Build()), nil
			},
		},
		di.Def{
			Name: constants.SERVICE,
			Build: func(ctn di.Container) (interface{}, error) {
				return services.NewService(builder.Build()), nil
			},
		},
		di.Def{
			Name: constants.LIBRARY,
			Build: func(ctn di.Container) (interface{}, error) {
				return libraries.NewLibrary(builder.Build()), nil
			},
		},
	)

	return builder.Build()
}
