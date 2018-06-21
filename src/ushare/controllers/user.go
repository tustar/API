package controllers

import (
	"github.com/gin-gonic/gin"
	"ushare/db"
	"net/http"
	"ushare/helpers"
	"strconv"
	"ushare/models"
	"ushare/middlewares"
	"log"
	"ushare/logger"
)

func UserCode(c *gin.Context) {
	user := new(db.User)
	user.Mobile = c.Request.FormValue("mobile")
	user.Captcha = helpers.GenerateCaptcha()
	user.Nick = "大神" + user.Captcha

	captcha := new(models.Captcha)
	captcha.Value = user.Captcha

	if id, captcha, err := user.Insert(); err != nil {
		c.JSON(http.StatusExpectationFailed, models.Result{
			Code:    helpers.Failure,
			Message: err.Error(),
			Data:    "",
			Extra:   "",
		})
		log.Fatal(err)
	} else {
		user.ID = int64(id)
		c.JSON(http.StatusOK, models.Result{
			Code:    helpers.OK,
			Message: "SUCCESS",
			Data:    models.Captcha{Value: captcha},
			Extra:   "",
		})
	}
}

func UserWeight(c *gin.Context) {
	mobile := c.Request.FormValue("mobile")
	weight, err := strconv.Atoi(c.Request.FormValue("weight"))
	if err != nil {
		c.JSON(http.StatusExpectationFailed, models.Result{
			Code:    helpers.Failure,
			Message: err.Error(),
			Data:    "",
			Extra:   "",
		})
		return
	}

	if user, err := db.Weight(mobile, weight); err != nil {
		c.JSON(http.StatusExpectationFailed, models.Result{
			Code:    helpers.Failure,
			Message: err.Error(),
			Data:    "",
			Extra:   "",
		})
	} else {
		c.JSON(http.StatusOK, models.Result{
			Code:    helpers.OK,
			Message: "SUCCESS",
			Data:    user,
			Extra:   "",
		})
	}
}

func UserLogin(c *gin.Context) {
	user := new(db.User)
	user.Mobile = c.Request.FormValue("mobile")
	user.Captcha = c.Request.FormValue("captcha")

	u, err := db.OneUserByMobile(user.Mobile)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, models.Result{
			Code:    helpers.Failure,
			Message: helpers.InsertFail,
			Data:    "",
			Extra:   "",
		})
		logger.D(err)
		return
	}

	if user.Captcha == u.Captcha {
		user.Token, err = middlewares.GenerateToken(user)
		c.JSON(http.StatusOK, models.Result{
			Code:    helpers.OK,
			Message: helpers.MsgSuccess,
			Data:    user,
			Extra:   "",
		})
	} else {
		c.JSON(http.StatusOK, models.Result{
			Code:    helpers.OK,
			Message: helpers.InvalidMsgCode,
			Data:    "",
			Extra:   "",
		})
	}
}

func UserList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Request.FormValue("page"))
	pageSize, _ := strconv.Atoi(c.Request.FormValue("page_size"))

	page = helpers.Max(page, 1)

	if pageSize == 0 {
		pageSize = 10
	}

	users, err := db.ListUser(page, pageSize)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, models.Result{
			Code:    helpers.Failure,
			Message: err.Error(),
			Data:    "",
			Extra:   "",
		})
	} else {
		c.JSON(http.StatusOK, models.Result{
			Code:    helpers.OK,
			Message: helpers.MsgSuccess,
			Data:    users,
			Extra:   "",
		})
	}
}
