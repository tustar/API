package models

import (
	"strconv"
	"ushare/db"
	"log"
)

type Topic struct {
	Id        int    `json:"id" form:"id"`
	UserId    string `json:"user_id" form:"user_id"`
	Title     string `json:"title" form:"title"`
	Desc      string `json:"description" form:"description"`
	CreatedAt int64  `json:"created_at" form:"created_at"`
	UpdatedAT int64  `json:"updated_at" form:"updated_at"`
	Begin     int64  `json:"begin" form:"begin"`
	End       int64  `json:"end" form:"end"`
	Shared    bool   `json:"shared" form:"shared"`
}

func ListTopic(page, pageSize int) (topics []Topic, err error) {
	topics = make([]Topic, 0)
	where := "WHERE 1=1"
	limit := strconv.Itoa((page-1)*pageSize) + "," + strconv.Itoa(pageSize)
	sql := "SELECT id, user_id, title, description, created_at, updated_at, begin, end, shared FROM topic " + where + " LIMIT " + limit
	log.Println(sql)
	rows, err := db.Conns.Query(sql)
	defer rows.Close()

	if err != nil {
		return
	}
	for rows.Next() {
		var topic Topic
		rows.Scan(&topic.Id, &topic.UserId, &topic.Title, &topic.Desc, &topic.CreatedAt, &topic.UpdatedAT, &topic.Begin,
			&topic.End, &topic.Shared)
		topics = append(topics, topic)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
