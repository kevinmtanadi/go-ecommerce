package services

import (
	"go-backend-template/src/constants"

	"github.com/sarulabs/di"
)

type BaseService struct {
	ioc     di.Container
	Service *Service
}

func NewBaseService(ioc di.Container) *BaseService {
	return &BaseService{
		ioc: ioc,
	}
}

func (s *BaseService) WithService() *Service {
	return s.ioc.Get(constants.SERVICE).(*Service)
}

type Service struct {
	Template TemplateService
	MySQL    MySQLService
	User     UserService
	Store    StoreService
}

func NewService(ioc di.Container) *Service {
	return &Service{
		Template: NewTemplateService(ioc),
		MySQL:    NewMySQLService(ioc),
		User:     NewUserService(ioc),
		Store:    NewStoreService(ioc),
	}
}
