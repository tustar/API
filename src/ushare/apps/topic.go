package apps

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"ushare/util"
	"net/http"
	"ushare/models"
)

func TopicList(c *gin.Context) {

	page, _ := strconv.Atoi(c.Request.FormValue("page"))
	pageSize, _ := strconv.Atoi(c.Request.FormValue("page_size"))

	page = util.Max(page, 1)

	if pageSize == 0 {
		pageSize = 10
	}

	topics, err := models.ListTopic(page, pageSize)
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
			"data":    topics,
			"extra":   "",
		})
	}
}
