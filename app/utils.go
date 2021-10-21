// Package app
// @Time  : 2021/10/19 18:44  
// @Author: mouse_zy@qq.com 
// @note  : 工具函数
package app

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"gorm.io/gorm"
	"net/http"
)

// HttpStatusCode 响应状态码类型
type HttpStatusCode uint16

// 全局变量集合
var (
	GStore = base64Captcha.DefaultMemStore // 全局验证码储存变量
	GDB    *gorm.DB                        // 全局数据库对象储存变量
)

// 全局状态码常量集合
const (
	HttpCodeOK      HttpStatusCode = 200   // 成功时状态码
	HttpCaptchaErr  HttpStatusCode = 40001 // 验证码为空或者验证码不正确
	HttpRegisterErr HttpStatusCode = 40002 // 注册参数获取失败
	HttpLoginErr    HttpStatusCode = 40003 // 登录参数获取失败
	HttpParamNil    HttpStatusCode = 40004 // 请求的JSON参数为空
	HttpPasswordErr HttpStatusCode = 40005 // 请求的密码不标准
	HttpDBErr       HttpStatusCode = 40006  // 查询数据库错误
	HttpMD5Err         HttpStatusCode = 40007  //加密模块错误
)

func initRequestJson(context *gin.Context, statusCode HttpStatusCode, data interface{}) {
	context.JSON(http.StatusOK, gin.H{
		"StatusCode": statusCode,
		"data":       data,
	})
}

// Err 异常时--响应的内容
func Err(context *gin.Context, statusCode HttpStatusCode, data interface{}) {
	initRequestJson(context, statusCode, data)
}

// Ok 正确时--响应内容
func Ok(context *gin.Context, data interface{}) {
	initRequestJson(context, HttpCodeOK, data)
}

func MD5(password string) (string, error) {
	// 1.0 加盐
	salt := []byte("zymouse")
	// 创建Hash
	h := md5.New()
	_, err := h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(salt)), err

}
