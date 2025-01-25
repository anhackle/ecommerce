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

// User Login documentation
// @Summary      List products at main page
// @Description  List products at main page by page & pagesize
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        page      query  int     true   "Page number"
// @Param        pagesize      query  int     true   "Number of products per page"
// @Success      200  {object}  response.ResponseData
// @Failure		 400  {object}  response.ResponseData
// @Failure		 401  {object}  response.ResponseData
// @Failure      500  {object}  response.ResponseData
// @Router       /products/ [get]
func (pc *ProductController) ListProductMainPage(c *gin.Context) {
	var input model.ProductsMainPageInput
	if err := c.ShouldBindQuery(&input); err != nil {
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
