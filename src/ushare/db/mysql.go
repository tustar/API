package db

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"ushare/config"
	"log"
)

var Conns *sql.DB

func init() {
	var err error
	username := config.Conf.Read("mysql", "username")
	password := config.Conf.Read("mysql", "password")
	dataname := config.Conf.Read("mysql", "dataname")
	port := config.Conf.Read("mysql", "port")
	host := config.Conf.Read("mysql", "host")

	dns := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dataname + "?parseTime=true"
	// fmt.Println(dns)
	Conns, err = sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = Conns.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	Conns.SetMaxIdleConns(20)
	Conns.SetMaxOpenConns(20)
}
