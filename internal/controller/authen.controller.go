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

// User Register documentation
// @Summary      Register API
// @Description  Register user with Email & Password
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        payload body model.UserInput true "payload"
// @Success      200  {object}  response.ResponseData	"Register success"
// @Failure		 400  {object}  response.ResponseData
// @Failure      500  {object}  response.ResponseData
// @Router       /authen/register [post]
func (ac *AuthenController) Register(c *gin.Context) {
	var input model.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, _ := ac.authenService.Register(c, input)

	response.HandleResult(c, result, nil)
}

// User Login documentation
// @Summary      Login API
// @Description  Login with Email & Password and return JWT
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        payload body model.UserInput true "payload"
// @Success      200  {object}  response.ResponseData	"Login success"
// @Failure		 400  {object}  response.ResponseData
// @Failure		 401  {object}  response.ResponseData
// @Failure      500  {object}  response.ResponseData
// @Router       /authen/login [post]
func (ac *AuthenController) Login(c *gin.Context) {
	var input model.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, jwtToken, _ := ac.authenService.Login(c, input)

	response.HandleResult(c, result, jwtToken)

}

func NewAuthenController(authenService service.IAuthenService) *AuthenController {
	return &AuthenController{
		authenService: authenService,
	}
}
