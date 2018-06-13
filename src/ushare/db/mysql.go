package db

import (
	_ "github.com/go-sql-driver/mysql"
	"ushare/config"
	"log"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	username := config.Conf.Read("mysql", "username")
	password := config.Conf.Read("mysql", "password")
	database := config.Conf.Read("mysql", "database")
	port := config.Conf.Read("mysql", "port")
	host := config.Conf.Read("mysql", "host")

	dns := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	// fmt.Println(dns)
	Db, err := gorm.Open("mysql", dns)

	if err != nil {
		log.Fatal(err.Error())
	}

	db := Db.DB()
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
}
