package service

import (
	"context"

	"github.com/anle/codebase/internal/model"
	"github.com/anle/codebase/internal/repo"
	"github.com/anle/codebase/internal/utils/hash"
	"github.com/anle/codebase/internal/utils/token"
	"github.com/anle/codebase/response"
	"gorm.io/gorm"
)

type IAuthenService interface {
	Register(ctx context.Context, input model.UserInput) (result int, err error)
	Login(ctx context.Context, input model.UserInput) (result int, jwtToken string, err error)
}

type authenService struct {
	authenRepo repo.IAuthenRepo
}

func (as *authenService) Register(ctx context.Context, input model.UserInput) (result int, err error) {
	_, err = as.authenRepo.FindByEmail(ctx, input)
	if err == nil {
		return response.ErrCodeUserHasExists, nil
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return response.ErrCodeInternal, err
	}

	hashPassword, err := hash.Hash(input.Password)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	var userInput = model.UserInput{
		Email:    input.Email,
		Password: hashPassword,
	}
	err = as.authenRepo.CreateUser(ctx, userInput)
	if err != nil {
		return response.ErrCodeInternal, err
	}

	return response.ErrCodeSuccess, nil
}

func (as *authenService) Login(ctx context.Context, input model.UserInput) (result int, jwtToken string, err error) {
	user, err := as.authenRepo.FindByEmail(ctx, input)
	if err != nil && err != gorm.ErrRecordNotFound {
		return response.ErrCodeInternal, "", err
	}

	if err != nil && err == gorm.ErrRecordNotFound {
		return response.ErrCodeLoginFail, "", err
	}

	if err := hash.Compare(user.Password, input.Password); err != nil {
		return response.ErrCodeLoginFail, "", err
	}

	jwtToken, err = token.GenJWTToken(user.ID)
	if err != nil {
		return response.ErrCodeInternal, "", err
	}

	return response.ErrCodeSuccess, jwtToken, nil
}

func NewAuthenService(authenRepo repo.IAuthenRepo) IAuthenService {
	return &authenService{
		authenRepo: authenRepo,
	}
}
