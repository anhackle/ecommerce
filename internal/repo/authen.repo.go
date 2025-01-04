package repo

import (
	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/internal/po"
)

type IAuthenRepo interface {
	CreateUser(userInput model.UserInput) (err error)
	FindByEmail(input model.UserInput) (user po.User, err error)
}

type authenRepo struct{}

func (ur *authenRepo) CreateUser(input model.UserInput) (err error) {
	var userInput = &po.User{
		Email:    input.Email,
		Password: input.Password,
	}

	err = global.Mdb.Model(&po.User{}).Create(&userInput).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *authenRepo) FindByEmail(input model.UserInput) (user po.User, err error) {
	err = global.Mdb.Model(&po.User{}).Where("email = ?", input.Email).First(&user).Error
	if err != nil {
		return po.User{}, err
	}

	return user, nil
}

func NewAuthenRepo() IAuthenRepo {
	return &authenRepo{}
}
