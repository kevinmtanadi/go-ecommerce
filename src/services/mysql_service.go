package services

import (
	"fmt"
	"go-backend-template/src/constants"

	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

type MySQLService interface {
	RunQuery(query string) (result []map[string]interface{}, err error)
}

type MySQLServiceImpl struct {
	service *BaseService
	db      *gorm.DB
}

func NewMySQLService(ioc di.Container) MySQLService {
	return &MySQLServiceImpl{
		service: NewBaseService(ioc),
		db:      ioc.Get(constants.MYSQL).(*gorm.DB),
	}
}

func (s *MySQLServiceImpl) RunQuery(query string) (result []map[string]interface{}, err error) {
	rows, err := s.db.Raw(query).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var row map[string]interface{}
		if err := s.db.ScanRows(rows, &row); err != nil {
			return nil, err
		}
		result = append(result, row)
	}

	fmt.Println(result)

	return result, nil
}
