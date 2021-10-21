### 1.0 登录和注册接口

### 2.0 路由

```
1. 登录[POST]：/app/register
2. 注册[POST]：/app/login
3. 验证码[GET]：/app/base/code 
```

### 3.0 POST消息格式--JSON

##### 3.1 登录

```
源码：
{
  "captcha_digit": "string",
  "captcha_id": "string",
  "name": "zymouse",
  "password": "123456789"
}
解析：
{
  "captcha_digit": 验证码数值,
  "captcha_id": "验证码ID",
  "name": 用户名,
  "password": 密码,
}
```

##### 3.2 注册

```
源码：
{
  "captcha_digit": "string",
  "captcha_id": "string",
  "gender": true,
  "name": "string",
  "password": "string",
  "phone": 0,
  "repassword": "string"
}
解析：
{
  "captcha_digit": 验证码数值,
  "captcha_id": "验证码ID",
  "gender": true表示男,False表示女
  "name": 用户名,
  "password": 密码,
  "phone": 手机号,
  "repassword": 再次输入密码
}
```

### 4.0 接口测试方法

```
http://localhost:14250/swagger/index.html#/
```

![](.\doc\001.png)

