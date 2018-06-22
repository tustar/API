package controllers

import (
	"github.com/gin-gonic/gin"
	"ushare/db"
	"ushare/helpers"
	"ushare/models"
	"net/http"
	"log"
)

func UserCaptcha(c *gin.Context) {

	mobile := c.Request.FormValue("mobile")
	if code, err := db.InsertCaptcha(mobile); err != nil {
		c.JSON(http.StatusExpectationFailed, models.Result{
			Code:    helpers.Failed,
			Message: err.Error(),
			Data:    "",
			Extra:   "",
		})
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, models.Result{
			Code:    helpers.OK,
			Message: "SUCCESS",
			Data: struct {
				Code string `json:"code"`
			}{Code: code},
			Extra: "",
		})
	}
}
