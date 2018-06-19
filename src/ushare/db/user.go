package db

type User struct {
	BaseModel
	Mobile  string  `json:"mobile" form:"mobile" gorm:"not null;unique"`
	Captcha string  `json:"captcha" form:"captcha"`
	Weight  int     `json:"weight" form:"weight"`
	Shared  bool    `json:"shared" form:"shared"`
	Nick    string  `json:"nick" form:"nick"`
	Type    string  `json:"type" from:"type"`
	Token   string  `json:"token" from:"token" gorm:"-"`
	Topics  []Topic `gorm:"ForeignKey:UserID"`
}

func (user User) Insert() (id int64, captcha string, err error) {
	result := Conn.Where(User{Mobile: user.Mobile}).FirstOrCreate(&user)
	id = user.ID
	captcha = user.Captcha
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (user *User) Users() (users []User, err error) {
	if err = Conn.Find(&users).Error; err != nil {
		return
	}
	return
}

func OneUserByMobile(mobile string) (user User, err error) {
	result := Conn.Find(&user, "mobile = ?", mobile)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func ListUser(page, pageSize int) (users []User, err error) {

	users = make([]User, 0)
	result := Conn.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
