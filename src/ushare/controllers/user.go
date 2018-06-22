package controllers

import (
	"github.com/gin-gonic/gin"
	"ushare/db"
	"net/http"
	"ushare/helpers"
	"strconv"
	"ushare/models"
	"ushare/middlewares"
	"ushare/logger"
)

func UserWeight(c *gin.Context) {
	mobile := c.Request.FormValue("mobile")
	weight, err := strconv.Atoi(c.Request.FormValue("weight"))
	if err != nil {
		c.JSON(http.StatusExpectationFailed, models.Result{
			Code:    helpers.Failed,
			Message: err.Error(),
			Data:    "",
			Extra:   "",
		})
		return
	}

	if user, err := db.UpdateUserWeight(mobile, weight); err != nil {
		c.JSON(http.StatusExpectationFailed, models.Result{
			Code:    helpers.Failed,
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

func UserNick(c *gin.Context) {
	mobile := c.Request.FormValue("mobile")
	nick := c.Request.FormValue("nick")

	if user, err := db.UpdateUserNick(mobile, nick); err != nil {
		c.JSON(http.StatusExpectationFailed, models.Result{
			Code:    helpers.Failed,
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
	captcha := c.Request.FormValue("captcha")

	code, err := db.GetCode(user.Mobile)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, models.Result{
			Code:    helpers.Failed,
			Message: err.Error(),
			Data:    "",
			Extra:   "",
		})
		logger.D(err)
		return
	}

	if captcha == code {
		user.Type = "user"
		_, err := user.InsertUser()
		if err != nil {
			c.JSON(http.StatusExpectationFailed, models.Result{
				Code:    helpers.Failed,
				Message: helpers.InsertFailed,
				Data:    "",
				Extra:   "",
			})
			logger.D(err)
		} else {
			user.Token, err = middlewares.GenerateToken(user)
			c.JSON(http.StatusOK, models.Result{
				Code:    helpers.OK,
				Message: helpers.Success,
				Data:    user,
				Extra:   "",
			})
		}
	} else {
		c.JSON(http.StatusOK, models.Result{
			Code:    helpers.OK,
			Message: helpers.InvalidCaptcha,
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
			Code:    helpers.Failed,
			Message: err.Error(),
			Data:    "",
			Extra:   "",
		})
	} else {
		c.JSON(http.StatusOK, models.Result{
			Code:    helpers.OK,
			Message: helpers.Success,
			Data:    users,
			Extra:   "",
		})
	}
}

func UserInfo(c *gin.Context) {

	mobile := c.Request.FormValue("mobile")
	user, err := db.OneUserByMobile(mobile)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, models.Result{
			Code:    helpers.Failed,
			Message: err.Error(),
			Data:    "",
			Extra:   "",
		})
		logger.D(err)
	} else {
		c.JSON(http.StatusOK, models.Result{
			Code:    helpers.OK,
			Message: "SUCCESS",
			Data:    user,
			Extra:   "",
		})
	}
}
