package initialize

import "github.com/gin-gonic/gin"

func Run() *gin.Engine {
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRedis()
	InitValidator()

	r := InitRouter()

	return r

}
