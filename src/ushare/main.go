package main

import (
	"ushare/db"
	"github.com/gin-gonic/gin"
	"ushare/routers"
	"ushare/config"
)

func main() {
	defer db.Db.Close()
	gin.SetMode(gin.DebugMode)
	router := routers.InitRouter()
	router.Run(":" + config.Conf.Read("site", "port"))
}
