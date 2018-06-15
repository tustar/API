package main

import (
	"github.com/gin-gonic/gin"
	"ushare/routers"
	"ushare/config"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := routers.InitRouter()
	router.Run(":" + config.Conf.Site.Port)
}
