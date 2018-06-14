package main

import (
	"github.com/gin-gonic/gin"
	"ushare/routers"
	"ushare/config"
	"ushare/db"
	"ushare/models"
)

func main() {
	defer db.Db.Close()
	gin.SetMode(gin.DebugMode)
	router := routers.InitRouter()
	router.Run(":" + config.Conf.Site.Port)
}
