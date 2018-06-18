package main

import (
	"github.com/gin-gonic/gin"
	"ushare/routers"
	"ushare/config"
	"ushare/db"
	"strconv"
)

func main() {
	// database
	defer db.Conn.Close()
	gin.SetMode(config.GinMode)
	router := routers.InitRouter()
	router.Run(":" + strconv.Itoa(config.SitePort))
}
