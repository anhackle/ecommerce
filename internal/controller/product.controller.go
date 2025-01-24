package controller

import (
	"github.com/anle/codebase/internal/model"
	service "github.com/anle/codebase/internal/services"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService service.IProductService
}

func (pc *ProductController) ListProductMainPage(c *gin.Context) {
	var input model.ProductsMainPageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponseExternal(c, response.ErrCodeExternal, nil)
		return
	}

	result, output, _ := pc.productService.GetProductsForMainPage(c, input)

	response.HandleResult(c, result, output)
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}
