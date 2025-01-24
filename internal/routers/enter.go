package routers

import (
	"github.com/anle/codebase/internal/routers/authen"
	"github.com/anle/codebase/internal/routers/product"
)

type RouterGroup struct {
	Authen  authen.AuthenRouter
	Product product.ProductRouter
}

var RouterGroupApp = new(RouterGroup)
