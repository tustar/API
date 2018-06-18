package config

import (
	"github.com/gin-gonic/gin"
)

const (
	// common
	GinMode     = gin.DebugMode
	Debug       = true
	GormLogMode = Debug

	// site
	SiteName = "ushare"
	SitePort = 4000

	// mysql
	MySqlHost     = "127.0.0.1"
	MySqlPort     = 3306
	MySqlUsername = "root"
	MySqlPassword = "TuStar1030!"
	MySqlDatabase = "ushare"

	// token
	TokenKey = "tustar"
)
