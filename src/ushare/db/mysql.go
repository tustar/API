package db

import (
	_ "github.com/go-sql-driver/mysql"
	"ushare/config"
	"github.com/jinzhu/gorm"
	"log"
)

var Instance *gorm.DB

func init() {

	username := config.Conf.MySql.Username
	password := config.Conf.MySql.Password
	database := config.Conf.MySql.Database
	port := config.Conf.MySql.Port
	host := config.Conf.MySql.Host

	dns := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	// fmt.Println(dns)
	Conn, err := gorm.Open("mysql", dns)
	if err != nil {
		log.Fatalf("mysql connect error %v", err)
	}

	debug := config.Conf.Build.Debug
	Conn.LogMode(debug)

	if Conn.Error != nil {
		log.Fatalf("database error %v", Conn.Error)
	}
}
