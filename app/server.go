// Package app
// @Time  : 2021/10/19 18:30  
// @Author: mouse_zy@qq.com 
// @note  : 路由的功能函数
package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

// Register 注册
// @Summary 注册用户接口
// @Schemes
// @Accept application/json
// @Param register body app.UserRegisterInfo true "用户信息"
// @Success 200 {string} string "成功"
// @Failure 40002 {string} string "参数绑定错误"
// @Failure 40001 {string} string "账户与密码不能为空"
// @Router /register [POST]
func Register(context *gin.Context) {
	// 1.0 获取请求JSON参数
	var requestJson UserRegisterInfo
	err := context.ShouldBindJSON(&requestJson)
	if err != nil {
		Err(context, HttpRegisterErr, "获取请求JSON数据失败")
		return
	}
	// 2.  判断请求JSON参数不标准
	{
		// 2.0 用户名，密码，再次密码为空
		if requestJson.Name == "" || requestJson.Password == "" || requestJson.RePassword == "" {
			Err(context, HttpParamNil, "用户名或者密码为空")
			return
		}
		// 2.1 验证码为空
		// 2.2 密码和再次密码不一样
		if requestJson.Password != requestJson.RePassword {
			Err(context, HttpCaptchaErr, "两次密码不一样")
			return
		}
	}
	// 3.0 校验验证码
	//if !GStore.Verify(requestJson.CaptchaId, requestJson.CaptchaDigit, true) {
	//	Err(context, HttpCaptchaErr, "验证码不正确")
	//	return
	//}
	// 3.1 查询数据库 -- 用户存在
	exits := GDB.Where("name = ?", requestJson.Name).First(&UserRegisterInfo{}).Error
	if exits == nil {
		fmt.Println(exits)
		Err(context, HttpDBErr, "该用户已存在")
		return
	}
	// 3.2 密码加密
	MDPassword, err := MD5(requestJson.Password)
	if err != nil {
		Err(context, HttpMD5Err, "注册失败")
	}
	requestJson.Password = MDPassword
	// 4.0 插入数据库
	err = GDB.Create(&requestJson).Error
	if err != nil {
		Err(context, HttpDBErr, "插入数据失败")
		return
	}
	Ok(context, "注册成功")
}

// Login 登录
// @Summary 登陆用户接口
// @Schemes
// @Accept application/json
// @Param register body app.UserLoginInfo true "用户信息"
// @Success 200 {string} string "成功"
// @Failure 40003 {string} string "参数绑定错误"
// @Failure 40001 {string} string "账户与密码不能为空"
// @Router /login [POST]
func Login(context *gin.Context) {
	// 1.0 获取请求JSON参数
	var responseJson UserLoginInfo
	err := context.ShouldBindJSON(&responseJson)
	if err != nil {
		Err(context, HttpLoginErr, "获取请求JSON数据失败")
		return
	}
	// 2.  判断请求JSON参数不标准
	{
		// 2.0 用户名和密码为空
		if responseJson.Name == "" || responseJson.Password == "" {
			Err(context, HttpParamNil, "用户名或者密码为空")
			return
		}
		// 2.1 验证码为空
		if responseJson.CaptchaDigit == "" || responseJson.CaptchaId == "" {
			Err(context, HttpParamNil, "验证码为空")
			return
		}
	}
	// 3.0 校验验证码
	//if !GStore.Verify(responseJson.CaptchaId, responseJson.CaptchaDigit, false) {
	//	Err(context, HttpCaptchaErr, "验证码不正确")
	//	return
	//}
	// 4.
	// 4.1 密码加密
	MDPassword, err := MD5(responseJson.Password)
	if err != nil {
		Err(context, HttpMD5Err, "注册失败")
	}
	userInfo := UserRegisterInfo{}
	// 4.2 查询数据库-用户名和密码
	exits := GDB.Where("name = ? AND password = ?", responseJson.Name, MDPassword).First(&userInfo).Error
	//fmt.Println("DEBUG:", exits != nil)
	if exits != nil {
		Err(context, HttpDBErr, "密码和用户名不正确")
		return
	}
	// 5.0 登录成功--返回用户基本信息
	Ok(context, userInfo.UserBaseInfo)

}

// Code 验证码逻辑
// @Summary 验证码接口
// @Description 生成验证编码码和图片
// @schemes
// @Accept  json
// @Success 200 {string} string "成功"
// @Failure 40001 {string} string ""验证码生成错误""
// @Router /base/code [GET]
func Code(context *gin.Context) {
	// 1. 生成随机验证码和图片
	// 1.0 验证码照片规格
	driverDigit := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	// 1.1 保存验证码信息
	captcha := base64Captcha.NewCaptcha(driverDigit, GStore)
	// 1.2 读取验证码
	codeId, photo, err := captcha.Generate()
	if err != nil {
		Err(context, HttpCaptchaErr, "验证码生成错误")
		return
	}
	// 3.0 响应验证码图片
	Ok(context, gin.H{"codeId": codeId, "photo": photo})
}
