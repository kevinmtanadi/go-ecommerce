package libraries

import (
	"context"
	"go-backend-template/src/constants"

	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

type TemplateLibrary interface {
	TemplateFunc(ctx context.Context, params interface{}) (data interface{}, err error)
}

type TemplateLibraryImpl struct {
	db *gorm.DB
}

func NewTemplateLibrary(ioc di.Container) TemplateLibrary {
	return &TemplateLibraryImpl{
		db: ioc.Get(constants.MYSQL).(*gorm.DB),
	}
}

func (l *TemplateLibraryImpl) TemplateFunc(ctx context.Context, params interface{}) (data interface{}, err error) {
	return nil, nil
}
