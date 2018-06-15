package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"ushare/helpers"
	"strings"
	"ushare/models"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sign := c.Request.FormValue("sign")
		method := c.Request.Method

		switch method {
		case http.MethodPost, http.MethodPut:
			if !Sign(c.Request, sign) {
				noAuth(c, helpers.Unauthorized)
				return
			}
		default:

		}

		c.Next()

		return
	}
}

func noAuth(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, models.Result{
		Code:    helpers.Failure,
		Message: msg,
		Data:    "",
		Extra:   "",
	})
	c.Abort()
}

func Sign(request *http.Request, sign string) bool {
	request.ParseForm()
	params := make(map[string]interface{})
	for k, v := range request.Form {
		if strings.EqualFold(k, "sign") {
			continue
		}
		params[k] = v[0]
	}

	err := CheckSign(params, sign)
	if err != nil {
		return false
	}

	return true
}
