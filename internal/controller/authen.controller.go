package controller

import (
	"github.com/anle/codebase/internal/model"
	service "github.com/anle/codebase/internal/services"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

type AuthenController struct {
	authenService service.IAuthenService
}

func (ac *AuthenController) Register(c *gin.Context) {
	var input model.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, _ := ac.authenService.Register(input)

	response.HandleResult(c, result, nil)
}

func (ac *AuthenController) Login(c *gin.Context) {
	var input model.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, jwtToken, _ := ac.authenService.Login(input)

	response.HandleResult(c, result, jwtToken)

}

func NewAuthenController(authenService service.IAuthenService) *AuthenController {
	return &AuthenController{
		authenService: authenService,
	}
}