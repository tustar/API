package models

type Captcha struct {
	Value string `json:"captcha" form:"captcha"`
}
