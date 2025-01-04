package main

import (
	"github.com/anle/codebase/internal/initialize"

	_ "github.com/anle/codebase/cmd/swag/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Ecomerce API
// @version         1.0
// @description     Ecomerce API.
// @termsOfService  https://github.com/anhackle/ecommerce

// @contact.name   ITS Group Developer
// @contact.url    https://github.com/anhackle/ecommerce
// @contact.email  nguyencaothai.vn@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8082
// @BasePath  /v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	r := initialize.Run()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8082")
}
