package apps

import (
	"github.com/gin-gonic/gin"
	"ushare/models"
	"net/http"
	"log"
	"ushare/util"
	"strconv"
)

func UserCode(c *gin.Context) {
	user := new(models.User)
	user.Mobile = c.Request.FormValue("mobile")
	user.Code = util.GenerateCode()

	code := new(models.Code)
	code.VCode = user.Code

	if id, err := user.AddUser(); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  util.Failure,
			"message": err.Error(),
			"data":    "",
			"extra":   "",
		})
		log.Fatal(err)
	} else {
		user.Id = int(id)
		c.JSON(http.StatusOK, gin.H{
			"status":  util.OK,
			"message": "SUCCESS",
			"data":    code,
			"extra":   "",
		})
	}
}

func UserLogin(c *gin.Context) {
	user := new(models.User)
	user.Mobile = c.Request.FormValue("mobile")
	user.Code = c.Request.FormValue("code")

	u, err := models.OneUserByMobile(user.Mobile)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  util.Failure,
			"message": util.InsertFail,
			"data":    "",
			"extra":   "",
		})
		log.Println(err)
		return
	}

	if user.Code == u.Code {
		c.JSON(http.StatusOK, gin.H{
			"status":  util.OK,
			"message": util.MsgSuccess,
			"data":    user,
			"extra":   "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  util.OK,
			"message": util.InvalidMsgCode,
			"data":    "",
			"extra":   "",
		})
	}
}

func UserList(c *gin.Context) {

	page, _ := strconv.Atoi(c.Request.FormValue("page"))
	pageSize, _ := strconv.Atoi(c.Request.FormValue("page_size"))

	page = util.Max(page, 1)

	if pageSize == 0 {
		pageSize = 10
	}

	users, err := models.ListUser(page, pageSize)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  util.Failure,
			"message": err.Error(),
			"data":    "",
			"extra":   "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  util.OK,
			"message": util.MsgSuccess,
			"data":    users,
			"extra":   "",
		})
	}
}

func UserGet(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Param("id"))

	user, err := models.OneUserById(uid)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
			"extra":   "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "SUCCESS",
			"data":    user,
			"extra":   "",
		})
	}

}

func UserEdit(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Param("id"))
	user := new(models.User)

	user.Id = uid
	user.Mobile = c.Request.FormValue("mobile")
	user.Nick = c.Request.FormValue("nick")
	user.Weight, _ = strconv.Atoi(c.Request.FormValue("weight"))

	if _, err := user.UpdateUser(uid); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
			"extra":   "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": util.MsgSuccess,
			"data":    "",
			"extra":   "",
		})
	}
}

func UserDelete(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Param("id"))

	if _, err := models.DeleteUser(uid); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
			"extra":   "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": util.MsgSuccess,
			"data":    "",
			"extra":   "",
		})
	}
}

func UserNick(c *gin.Context) {

	mobile := c.Request.FormValue("mobile")
	nick := c.Request.FormValue("nick")

	if _, err := models.UpdateUserNick(nick, mobile); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
			"extra":   "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": util.MsgSuccess,
			"data":    "",
			"extra":   "",
		})
	}
}

func UserWeight(c *gin.Context) {

	mobile := c.Request.FormValue("mobile")
	weight, _ := strconv.Atoi(c.Request.FormValue("weight"))

	if _, err := models.UpdateUserWeight(weight, mobile); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
			"extra":   "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": util.MsgSuccess,
			"data":    "",
			"extra":   "",
		})
	}
}

func UserShared(c *gin.Context) {

	mobile := c.Request.FormValue("mobile")
	shared, _ := strconv.ParseBool(c.Request.FormValue("shared"))

	if _, err := models.UpdateUserShared(shared, mobile); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    "",
			"extra":   "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": util.MsgSuccess,
			"data":    "",
			"extra":   "",
		})
	}
}

