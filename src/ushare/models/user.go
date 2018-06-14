package models

import "ushare/db"

type User struct {
	BaseModel
	Mobile string  `json:"mobile" form:"mobile"`
	Code   string  `json:"code" form:"code"`
	Token  string  `json:"token" form:"token" gorm:"-"`
	Weight int     `json:"weight" form:"weight"`
	Shared bool    `json:"shared" form:"shared"`
	Nick   string  `json:"nick" form:"nick"`
	Type   string  `json:"type" from:"type"`
	Topics []Topic `gorm:"ForeignKey:UserID"`
}

func (user *User) Insert() (id uint, err error) {
	db.Db.Create(&user)
	if db.Db.Error != nil {
		return 0, db.Db.Error
	}
	return user.ID, nil
}
