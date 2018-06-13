package actions

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"ushare/helpers"
	"net/http"
	"ushare/models"
)

func TopicList(c *gin.Context) {

	page, _ := strconv.Atoi(c.Request.FormValue("page"))
	pageSize, _ := strconv.Atoi(c.Request.FormValue("page_size"))

	page = helpers.Max(page, 1)

	if pageSize == 0 {
		pageSize = 10
	}

	topics, err := models.ListTopic(page, pageSize)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"status":  helpers.Failure,
			"message": err.Error(),
			"data":    "",
			"extra":   "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  helpers.OK,
			"message": helpers.MsgSuccess,
			"data":    topics,
			"extra":   "",
		})
	}
}
