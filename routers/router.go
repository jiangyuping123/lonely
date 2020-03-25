package routers

import (
	"AdapterServer/conf"
	"AdapterServer/middleware/cors"
	"AdapterServer/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	appCfg := conf.GetAppCfg()
	gin.SetMode(appCfg.RunMode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.CorsHandler())
	r.Use(gin.Recovery())

	r.POST("/login", api.Login)
	r.GET("/ports", api.ListPorts)

	return r
}
