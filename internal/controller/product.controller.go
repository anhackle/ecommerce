package controller

import (
	"github.com/anle/codebase/internal/model"
	service "github.com/anle/codebase/internal/services"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productController service.IProductService
}

func (pc *ProductController) ListProductMainPage(c *gin.Context) {
	var input model.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, _ := pc.authenService.Register(c, input)

	response.HandleResult(c, result, nil)
}
