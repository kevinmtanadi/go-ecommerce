package libraries

import "github.com/sarulabs/di"

type Library struct {
	Template TemplateLibrary
	User     UserLibrary
}

func NewLibrary(ioc di.Container) *Library {
	return &Library{
		Template: NewTemplateLibrary(ioc),
		User:     NewUserLibrary(ioc),
	}
}
