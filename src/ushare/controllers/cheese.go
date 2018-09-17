package controllers

import (
	"github.com/gin-gonic/gin"
	"ushare/db"
	"net/http"
	"ushare/helpers"
	"strconv"
	"ushare/models"
)


func CheeseList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Request.FormValue("page"))
	pageSize, _ := strconv.Atoi(c.Request.FormValue("page_size"))

	page = helpers.Max(page, 1)

	if pageSize == 0 {
		pageSize = 10
	}

	users, err := db.ListCheese(page, pageSize)
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
