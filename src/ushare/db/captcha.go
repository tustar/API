package db

import (
	"time"
	"ushare/helpers"
)

type Captcha struct {
	BaseModel
	Mobile    string `json:"-" gorm:"not null;unique"`
	Code      string `json:"code" form:"code"`
	ExpiredAt time.Time
}

func InsertCaptcha(mobile string) (code string, err error) {
	captcha := new(Captcha)
	captcha.Code = helpers.GenerateCaptcha()
	captcha.ExpiredAt = time.Now().UTC().Add(3 * time.Hour)
	result := Conn.Where("mobile = ?", mobile).FirstOrCreate(&captcha)
	code = captcha.Code
	if err = result.Error; err != nil {
		return
	}
	return
}

func GetCode(mobile string) (code string, err error) {
	captcha := new(Captcha)
	result := Conn.Find(&captcha, "mobile = ?", mobile)
	code = captcha.Code
	if err = result.Error; err != nil {
		return
	}
	return
}
