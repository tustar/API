package db

import (
	_ "github.com/go-sql-driver/mysql"
	"ushare/config"
	"github.com/jinzhu/gorm"
	"ushare/logger"
	"fmt"
)

var Conn *gorm.DB

func init() {
	var err error
	username := config.MySqlUsername
	password := config.MySqlPassword
	database := config.MySqlDatabase
	port := config.MySqlPort
	host := config.MySqlHost

	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host,
		port, database)
	Conn, err = gorm.Open("mysql", dns)
	if err != nil {
		logger.E("Mysql open error ", err)
	}

	Conn.LogMode(config.GormLogMode)
	Conn.AutoMigrate(&User{}, &Topic{}, &Captcha{})

	db := Conn.DB()
	err = db.Ping()
	if err != nil {
		logger.E(err.Error())
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
}
