package db

import (
	_ "github.com/go-sql-driver/mysql"
	"ushare/config"
	"github.com/jinzhu/gorm"
	"ushare/models"
)

var Db *gorm.DB

func InitDB(*gorm.DB, error) {

	username := config.Conf.MySql.Username
	password := config.Conf.MySql.Password
	database := config.Conf.MySql.Database
	port := config.Conf.MySql.Port
	host := config.Conf.MySql.Host

	dns := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	// fmt.Println(dns)
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{}, &models.Topic{})
	debug := config.Conf.Build.Debug
	db.LogMode(debug)

	conn := db.DB()
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	conn.SetMaxIdleConns(20)
	conn.SetMaxOpenConns(20)

	return db, nil
}
