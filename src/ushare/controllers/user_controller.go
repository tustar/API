package controllers

import (
	"github.com/gin-gonic/gin"
	"ushare/models"
	"net/http"
	"log"
	"ushare/helpers"
)

type Code struct {
	VCode string `json:"v_code" form:"v_code"`
}

func UserCode(c *gin.Context) {
	user := new(models.User)
	user.Mobile = c.Request.FormValue("mobile")
	user.Code = helpers.GenerateCode()

	code := new(Code)
	code.VCode = user.Code

	if id, err := user.Insert(); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  helpers.Failure,
			"message": err.Error(),
			"data":    "",
			"extra":   "",
		})
		log.Fatal(err)
	} else {
		user.ID = uint(id)
		c.JSON(http.StatusOK, gin.H{
			"status":  helpers.OK,
			"message": "SUCCESS",
			"data":    code,
			"extra":   "",
		})
	}
}

//func UserLogin(c *gin.Context) {
//	user := new(models.User)
//	user.Mobile = c.Request.FormValue("mobile")
//	user.Code = c.Request.FormValue("code")
//
//	u, err := models.OneUserByMobile(user.Mobile)
//	if err != nil {
//		c.JSON(http.StatusExpectationFailed, gin.H{
//			"status":  helpers.Failure,
//			"message": helpers.InsertFail,
//			"data":    "",
//			"extra":   "",
//		})
//		log.Println(err)
//		return
//	}
//
//	if user.Code == u.Code {
//		c.JSON(http.StatusOK, gin.H{
//			"status":  helpers.OK,
//			"message": helpers.MsgSuccess,
//			"data":    user,
//			"extra":   "",
//		})
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"status":  helpers.OK,
//			"message": helpers.InvalidMsgCode,
//			"data":    "",
//			"extra":   "",
//		})
//	}
//}
//
//func UserList(c *gin.Context) {
//
//	page, _ := strconv.Atoi(c.Request.FormValue("page"))
//	pageSize, _ := strconv.Atoi(c.Request.FormValue("page_size"))
//
//	page = helpers.Max(page, 1)
//
//	if pageSize == 0 {
//		pageSize = 10
//	}
//
//	users, err := models.ListUser(page, pageSize)
//	if err != nil {
//		c.JSON(http.StatusExpectationFailed, gin.H{
//			"status":  helpers.Failure,
//			"message": err.Error(),
//			"data":    "",
//			"extra":   "",
//		})
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"status":  helpers.OK,
//			"message": helpers.MsgSuccess,
//			"data":    users,
//			"extra":   "",
//		})
//	}
//}
