package libraries

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"go-backend-template/src/constants"
	"go-backend-template/src/model"

	"github.com/sarulabs/di"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserLibrary interface {
	GetUserDataByEmail(ctx context.Context, email string) (data model.User, err error)
	GetUserDataByUsername(ctx context.Context, username string) (data model.User, err error)
	Create(ctx context.Context, params model.User) (err error)
	Update(ctx context.Context, params model.User) (err error)
	EncryptPassword(password string) (string, string, error)
}

type UserLibraryImpl struct {
	db *gorm.DB
}

func NewUserLibrary(ioc di.Container) UserLibrary {
	return &UserLibraryImpl{
		db: ioc.Get(constants.MYSQL).(*gorm.DB),
	}
}

func (l *UserLibraryImpl) GetUserDataByEmail(ctx context.Context, email string) (data model.User, err error) {
	err = l.db.WithContext(ctx).Where("email = ?", email).First(&data).Error
	return
}

func (l *UserLibraryImpl) GetUserDataByUsername(ctx context.Context, username string) (data model.User, err error) {
	err = l.db.WithContext(ctx).Where("username = ?", username).First(&data).Error
	return
}

func (l *UserLibraryImpl) Create(ctx context.Context, params model.User) (err error) {
	err = l.db.WithContext(ctx).Create(&params).Error
	return
}

func (l *UserLibraryImpl) EncryptPassword(password string) (string, string, error) {
	salt := generateSalt()

	byteString := []byte(password + salt)

	hashedPassword, err := bcrypt.GenerateFromPassword(byteString, bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	return string(hashedPassword), salt, nil
}

func (l *UserLibraryImpl) Update(ctx context.Context, params model.User) (err error) {
	err = l.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", params.ID).Updates(&params).Error
	return
}

func generateSalt() string {
	saltBytes := make([]byte, 16)
	rand.Read(saltBytes)

	return base64.RawURLEncoding.EncodeToString(saltBytes)
}
