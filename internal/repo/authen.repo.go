package repo

import (
	"context"

	"github.com/anle/codebase/global"
	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/internal/po"
)

type IAuthenRepo interface {
	CreateUser(ctx context.Context, userInput model.UserInput) (err error)
	FindByEmail(ctx context.Context, input model.UserInput) (user po.User, err error)
}

type authenRepo struct{}

func (ur *authenRepo) CreateUser(ctx context.Context, input model.UserInput) (err error) {
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

func (ur *authenRepo) FindByEmail(ctx context.Context, input model.UserInput) (user po.User, err error) {
	err = global.Mdb.Model(&po.User{}).Where("email = ?", input.Email).First(&user).Error
	if err != nil {
		return po.User{}, err
	}

	return user, nil
}

func NewAuthenRepo() IAuthenRepo {
	return &authenRepo{}
}
