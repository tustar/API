package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"ushare/helpers"
	"strings"
	"ushare/models"
	"ushare/logger"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sign := c.Request.FormValue("sign")
		token := c.Request.FormValue("token")
		method := c.Request.Method

		switch method {
		case http.MethodPost, http.MethodPut:
			if len(token) > 0 {
				valid, _ := ValidateToken(token)
				if !valid {
					logger.W("Token invalid")
					noAuth(c, helpers.ExpiredToken)
				}
			}
			if !Sign(c.Request, sign) {
				logger.W("Sign invalid")
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
		Code:    helpers.Failed,
		Message: msg,
		Data:    "",
		Extra:   "",
	})
	c.Abort()
}

func Sign(request *http.Request, sign string) bool {
	if sign == "" {
		logger.W("sign is empty")
		return false
	}

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
