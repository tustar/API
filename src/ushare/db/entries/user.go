package entries

import (
	"ushare/db"
)

type User struct {
	BaseModel
	Mobile  string  `json:"mobile" form:"mobile"`
	Captcha string  `json:"captcha" form:"captcha"`
	Weight  int     `json:"weight" form:"weight"`
	Shared  bool    `json:"shared" form:"shared"`
	Nick    string  `json:"nick" form:"nick"`
	Type    string  `json:"type" from:"type"`
	Topics  []Topic `gorm:"ForeignKey:UserID"`
}

func (user User) Insert() (id int64, err error) {
	result := db.Instance.Create(&user)
	id = user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (user *User) Users() (users []User, err error) {
	if err = db.Instance.Find(&users).Error; err != nil {
		return
	}
	return
}

func OneUserByMobile(mobile string) (user User, err error) {
	result := db.Instance.Find(&user, "mobile = ?", mobile)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func ListUser(page, pageSize int) (users []User, err error) {
	result := db.Instance.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
