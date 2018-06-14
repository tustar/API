package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ushare/middlewares"
	"ushare/controllers"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.StaticFS("/public", http.Dir("public"))
	//router.LoadHTMLGlob("templates/*")
	v1 := router.Group("/v1")
	v1.Use(middlewares.Auth())
	{
		/**
		 * User
		 */
		v1.POST("/user/code", controllers.UserCode)
		//v1.POST("/user/login", controllers.UserLogin)
		//// curl -X GET http://127.0.0.1:4000/v1/user
		//v1.GET("/user", controllers.UserList)
	}

	return router
}
