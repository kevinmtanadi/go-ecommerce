package libraries

import (
	"context"
	"go-backend-template/src/constants"
	"go-backend-template/src/model"

	"github.com/sarulabs/di"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductLibrary interface {
	Create(ctx context.Context, params model.Product) (err error)
	Update(ctx context.Context, params model.Product) (err error)
	Delete(ctx context.Context, productID int) (err error)
	UpsertAdditionalData(ctx context.Context, params model.ProductAdditionalData) (err error)
	RateProduct(ctx context.Context, params model.ProductRating) (err error)
}

type ProductLibraryImpl struct {
	db *gorm.DB
}

func NewProductLibrary(ioc di.Container) ProductLibrary {
	return &ProductLibraryImpl{
		db: ioc.Get(constants.MYSQL).(*gorm.DB),
	}
}

func (l *ProductLibraryImpl) Create(ctx context.Context, params model.Product) (err error) {
	err = l.db.WithContext(ctx).Create(&params).Error
	return
}

func (l *ProductLibraryImpl) Update(ctx context.Context, params model.Product) (err error) {
	err = l.db.WithContext(ctx).Model(&model.Product{}).Where("id = ?", params.ID).Updates(&params).Error
	return
}

func (l *ProductLibraryImpl) Delete(ctx context.Context, productID int) (err error) {
	err = l.db.WithContext(ctx).Delete(&model.Product{}, productID).Error
	return
}

func (l *ProductLibraryImpl) UpsertAdditionalData(ctx context.Context, params model.ProductAdditionalData) (err error) {
	err = l.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "product_id"}, {Name: "key"}},
		DoUpdates: clause.AssignmentColumns([]string{"value"}),
	}).Model(&model.Product{}).Where("id = ?", params.ID).Create(&params).Error

	return
}

func (l *ProductLibraryImpl) RateProduct(ctx context.Context, params model.ProductRating) (err error) {
	err = l.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "product_id"}, {Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"rating"}),
	}).Create(&params).Error
	return
}
