package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ushare/middlewares"
	"ushare/apps"
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
		v1.POST("/user/code", apps.UserCode)
		v1.POST("/user/login", apps.UserLogin)
		// curl -X GET http://127.0.0.1:4000/v1/user
		v1.GET("/user", apps.UserList)
		v1.GET("/user/:id", apps.UserGet)
		v1.PUT("/user/:id", apps.UserEdit)
		v1.POST("/user/nick", apps.UserNick)
		v1.POST("/user/weight", apps.UserWeight)
		v1.POST("/user/shared", apps.UserShared)
		v1.DELETE("/user/:id", apps.UserDelete)

		/**
		 * Topic
		 */
		v1.GET("/topic", apps.TopicList)
	}

	return router
}
