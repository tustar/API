package models

import (
	"ushare/db"
	"log"
	"strconv"
	"ushare/helpers"
)

type User struct {
	Id     int    `json:"id" form:"id"`
	Mobile string `json:"mobile" form:"mobile"`
	Code   string `json:"code" form:"code"`
	Token  string `json:"token" form:"token"`
	Weight int    `json:"weight" form:"weight"`
	Shared bool   `json:"shared" form:"shared"`
	Nick   string `json:"nick" form:"nick"`
	Type   string `json:"type" from:"type"`
	LastAt int64  `json:"last_at" form:"last_at"`
	NextAt int64  `json:"next_at" form:"next_at"`
}

func (user *User) AddUser() (id int64, err error) {
	res, err := db.Conns.Exec("INSERT INTO user(mobile, code) VALUES (?, ?)", user.Mobile, user.Code)
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	return
}

func (user *User) ReplaceUser() (id int64, err error) {
	res, err := db.Conns.Exec("REPLACE INTO user(mobile, code) VALUES (?, ?)",
		user.Mobile, user.Code)
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	return
}

func OneUserByMobile(mobile string) (user User, err error) {
	user.Id = 0
	user.Mobile = ""
	user.Code = ""
	err = db.Conns.QueryRow("SELECT id, mobile, code FROM user WHERE mobile=? LIMIT 1", mobile).Scan(
		&user.Id, &user.Mobile, &user.Code)
	return
}

func OneUserById(id int) (user User, err error) {
	user.Id = 0
	user.Mobile = ""
	user.Code = ""
	err = db.Conns.QueryRow("SELECT id, mobile, code FROM user WHERE id=? LIMIT 1", id).Scan(
		&user.Id, &user.Mobile, &user.Code)
	return
}

func (user *User) UpdateUser(id int) (count int64, err error) {
	res, err := db.Conns.Prepare("UPDATE user SET mobile=?,code=?, token=?, weight=?, shared=?, nick=? WHERE id=?")
	defer res.Close()
	if err != nil {
		log.Println(err)
	}
	rs, err := res.Exec(user.Mobile, user.Code, user.Token, user.Weight, user.Shared, user.Nick, user.Id)
	if err != nil {
		log.Println(err)
	}
	count, err = rs.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	return
}

func ListUser(page, pageSize int) (users []User, err error) {
	users = make([]User, 0)
	where := "WHERE 1=1"
	limit := strconv.Itoa((page-1)*pageSize) + "," + strconv.Itoa(pageSize)
	sql := "SELECT mobile, code, token, weight, shared, nick, type FROM user " + where + " LIMIT " + limit
	log.Println(sql)
	rows, err := db.Conns.Query(sql)
	defer rows.Close()

	if err != nil {
		return
	}
	for rows.Next() {
		var user User
		rows.Scan(&user.Mobile, &user.Code, &user.Token, &user.Weight, &user.Shared, &user.Nick, &user.Type)
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func DeleteUser(id int) (count int64, err error) {
	count = 0
	rs, err := db.Conns.Exec("DELETE FROM user WHERE id=?", id)
	if err != nil {
		log.Fatalln(err)
		return
	}
	count, err = rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
		return
	}
	return
}

func UpdateUserNick(nick string, mobile string) (count int64, err error) {
	res, err := db.Conns.Prepare("UPDATE user SET nick=? WHERE mobile=?")
	defer res.Close()
	if err != nil {
		log.Println(err)
	}
	rs, err := res.Exec(nick, mobile)
	if err != nil {
		log.Println(err)
	}
	count, err = rs.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	return
}

func UpdateUserWeight(weight int, mobile string) (count int64, err error) {
	res, err := db.Conns.Prepare("UPDATE user SET weight=? WHERE mobile=?")
	defer res.Close()
	if err != nil {
		log.Println(err)
	}
	rs, err := res.Exec(weight, mobile)
	if err != nil {
		log.Println(err)
	}
	count, err = rs.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	return
}

func UpdateUserShared(shared bool, mobile string) (count int64, err error) {
	res, err := db.Conns.Prepare("UPDATE user SET shared=? WHERE mobile=?")
	defer res.Close()
	if err != nil {
		log.Println(err)
	}
	rs, err := res.Exec(helpers.BoolToInt(shared), mobile)
	if err != nil {
		log.Println(err)
	}
	count, err = rs.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	return
}
