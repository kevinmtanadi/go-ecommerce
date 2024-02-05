package services

import (
	"context"
	"errors"
	"go-backend-template/src/constants"
	"go-backend-template/src/dtos"
	"go-backend-template/src/libraries"
	"go-backend-template/src/model"

	"github.com/sarulabs/di"
)

type StoreService interface {
	CreateNewStore(ctx context.Context, params dtos.CreateStoreReq) (data dtos.CreateStoreRes, err error)
	UpdateData(ctx context.Context, params dtos.UpdateStoreReq) (data dtos.UpdateStoreRes, err error)
	Delete(ctx context.Context, storeID int) (err error)
}

type StoreServiceImpl struct {
	service *BaseService
	library *libraries.Library
}

func NewStoreService(ioc di.Container) StoreService {
	return &StoreServiceImpl{
		service: NewBaseService(ioc),
		library: ioc.Get(constants.LIBRARY).(*libraries.Library),
	}
}

func (s *StoreServiceImpl) CreateNewStore(ctx context.Context, params dtos.CreateStoreReq) (data dtos.CreateStoreRes, err error) {
	store, err := s.library.Store.GetStoreDataByStoreName(ctx, params.Storename)
	if store.ID > 0 {
		return data, errors.New("store name already exists")
	}

	newStore := model.Store{
		OwnerID:     uint(params.UserID),
		StoreName:   params.Storename,
		Description: params.Description,
	}

	err = s.library.Store.Create(ctx, newStore)

	return
}

func (s *StoreServiceImpl) UpdateData(ctx context.Context, params dtos.UpdateStoreReq) (data dtos.UpdateStoreRes, err error) {
	store, err := s.library.Store.GetStoreDataByStoreName(ctx, params.Storename)
	if store.ID > 0 {
		return data, errors.New("store name already exists")
	}

	s.library.Store.Update(ctx, model.Store{
		ID:          uint(params.StoreID),
		StoreName:   params.Storename,
		Description: params.Description,
	})

	return
}

func (s *StoreServiceImpl) Delete(ctx context.Context, storeID int) (err error) {
	err = s.library.Store.Delete(ctx, storeID)
	return
}
