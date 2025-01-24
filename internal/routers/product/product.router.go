package product

import (
	"github.com/anle/codebase/internal/wire"
	"github.com/gin-gonic/gin"
)

type ProductRouter struct{}

func (ar *ProductRouter) InitProductRouter(router *gin.RouterGroup) {
	//public router
	productController, _ := wire.InitProductRouterHandler()

	productRouterPublic := router.Group("/product")

	{
		productRouterPublic.POST("/get_list_products", productController.ListProductMainPage)
	}
}
