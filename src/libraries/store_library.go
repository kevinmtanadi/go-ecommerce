package libraries

import (
	"context"
	"go-backend-template/src/constants"
	"go-backend-template/src/model"

	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

type StoreLibrary interface {
	GetStoreDataByStoreName(ctx context.Context, storeName string) (data model.Store, err error)
	GetStoreDataByStoreID(ctx context.Context, storeID int) (data model.Store, err error)
	Create(ctx context.Context, params model.Store) (err error)
	Update(ctx context.Context, params model.Store) (err error)
	Delete(ctx context.Context, storeID int) (err error)
}

type StoreLibraryImpl struct {
	db *gorm.DB
}

func NewStoreLibrary(ioc di.Container) StoreLibrary {
	return &StoreLibraryImpl{
		db: ioc.Get(constants.MYSQL).(*gorm.DB),
	}
}

func (l *StoreLibraryImpl) GetStoreDataByStoreName(ctx context.Context, storeName string) (data model.Store, err error) {
	err = l.db.WithContext(ctx).Where("store_name = ?", storeName).First(&data).Error
	return
}

func (l *StoreLibraryImpl) GetStoreDataByStoreID(ctx context.Context, storeID int) (data model.Store, err error) {
	err = l.db.WithContext(ctx).Where("id = ?", storeID).First(&data).Error
	return
}

func (l *StoreLibraryImpl) Create(ctx context.Context, params model.Store) (err error) {
	err = l.db.WithContext(ctx).Create(&params).Error
	return
}

func (l *StoreLibraryImpl) Update(ctx context.Context, params model.Store) (err error) {
	err = l.db.WithContext(ctx).Model(&model.Store{}).Where("id = ?", params.ID).Updates(&params).Error
	return
}

func (l *StoreLibraryImpl) Delete(ctx context.Context, storeID int) (err error) {
	err = l.db.WithContext(ctx).Delete(&model.Store{}, storeID).Error
	return
}
