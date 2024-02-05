package services

import (
	"context"
	"go-backend-template/src/dtos"

	"github.com/sarulabs/di"
)

type TemplateService interface {
	TemplateFunc(ctx context.Context, params dtos.TemplateDtoReq) (data dtos.TemplateDtoRes, err error)
}

type TemplateServiceImpl struct {
	service *BaseService
}

func NewTemplateService(ioc di.Container) TemplateService {
	return &TemplateServiceImpl{
		service: NewBaseService(ioc),
	}
}

func (s *TemplateServiceImpl) TemplateFunc(ctx context.Context, params dtos.TemplateDtoReq) (data dtos.TemplateDtoRes, err error) {
	return dtos.TemplateDtoRes{}, nil
}
