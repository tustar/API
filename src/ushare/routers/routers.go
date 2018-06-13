package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ushare/middlewares"
	"ushare/actions"
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
		v1.POST("/user/code", actions.UserCode)
		v1.POST("/user/login", actions.UserLogin)
		// curl -X GET http://127.0.0.1:4000/v1/user
		v1.GET("/user", actions.UserList)
		v1.GET("/user/:id", actions.UserGet)
		v1.PUT("/user/:id", actions.UserEdit)
		v1.POST("/user/nick", actions.UserNick)
		v1.POST("/user/weight", actions.UserWeight)
		v1.POST("/user/shared", actions.UserShared)
		v1.DELETE("/user/:id", actions.UserDelete)

		/**
		 * Topic
		 */
		v1.GET("/topic", actions.TopicList)
	}

	return router
}
