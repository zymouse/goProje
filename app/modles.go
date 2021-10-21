// Package app
// @Time  : 2021/10/19 20:10  
// @Author: mouse_zy@qq.com 
// @note  : sqLite3
package app

type UserBaseInfo struct {
	Name   string `json:"name" gorm:"column:name;comment:名字"`
	Gender bool   `json:"gender" gorm:"column:gender;comment:性别-真为男，假为女"`
	Phone  uint64 `json:"phone" gorm:"column:phone;comment:电话号码"`
}

type UserRegisterInfo struct {
	ID           uint   `json:"-" gorm:"primaryKey"`
	RePassword   string `json:"repassword" gorm:"-"`
	Password     string `json:"password" gorm:"column:password;comment:密码"`
	CaptchaDigit string `json:"captcha_digit" gorm:"-"`
	CaptchaId    string `json:"captcha_id" gorm:"-"`
	UserBaseInfo
}

type UserLoginInfo struct {
	Name         string `json:"name"`
	CaptchaDigit string `json:"captcha_digit"`
	CaptchaId    string	`json:"captcha_id"`
	Password     string `json:"password"`
}
