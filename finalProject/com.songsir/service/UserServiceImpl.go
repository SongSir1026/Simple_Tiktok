package service

import (
	"encoding/json"
	"finalProject/com.songsir/common"
	"finalProject/com.songsir/common/constant"
	"finalProject/com.songsir/common/utils"
	"finalProject/com.songsir/dao"
	"finalProject/com.songsir/dto"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"regexp"
	"time"
)

//用户注册
func Register(username string, password string) dto.LoginResponse {
	//防止Sql注入
	//const SQL_JUDGE = `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	m1, err1 := regexp.MatchString(constant.SQL_JUDGE, username)

	m2, err2 := regexp.MatchString(constant.SQL_JUDGE, password)
	if m1 || m2 {
		return dto.LoginResponse{StatusMsg: "无效账号或者密码", StatusCode: -1}
	}
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
		fmt.Println(err2)
	}
	//参数检验
	if username == "" || len(username) > 32 {
		return dto.LoginResponse{StatusMsg: "无效账号", StatusCode: -1}
	} else if password == "" || len(password) > 32 {
		return dto.LoginResponse{StatusMsg: "无效密码", StatusCode: -1}
	}

	userId := time.Now().Unix()
	pw := utils.MD5(password)
	var result = dao.Register(username, pw, userId)
	token := utils.RandString(32)
	if result <= 0 {
		return dto.LoginResponse{StatusMsg: "注册失败", StatusCode: -1}
	}
	userJson, err := json.Marshal(common.User{UserId: int(userId), Username: username, Password: pw})
	if err != nil {
		fmt.Println(err)
	}
	utils.Set(constant.USER_FLAG+token, userJson)
	utils.Expire(constant.USER_FLAG+token, 30)
	return dto.LoginResponse{StatusMsg: "注册成功", StatusCode: 0, UserId: int(userId), Token: token}
}

//登录接口
func Login(username string, password string) dto.LoginResponse {
	//防止Sql注入
	m1, err1 := regexp.MatchString(constant.SQL_JUDGE, username)
	m2, err2 := regexp.MatchString(constant.SQL_JUDGE, password)
	if m1 || m2 {
		return dto.LoginResponse{StatusMsg: "无效账号或者密码", StatusCode: -1}
	}
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
		fmt.Println(err2)
	}
	//参数检验
	if username == "" || len(username) > 32 {
		return dto.LoginResponse{StatusCode: -1, StatusMsg: "无效账号"}
	} else if password == "" || len(password) > 32 {
		return dto.LoginResponse{StatusCode: -1, StatusMsg: "无效密码"}
	}

	user := dao.Login(username)
	pw := utils.MD5(password)
	if user.Username != username {
		return dto.LoginResponse{StatusCode: -1, StatusMsg: "用户不存在"}
	} else if pw != user.Password {
		return dto.LoginResponse{StatusCode: -1, StatusMsg: "密码不正确"}
	}

	//获取redis连接
	conn := utils.GetRedisConnection()
	if conn == nil {
		fmt.Println("获取redis连接失败")
	}
	//将token存入用户信息存入redis中
	token := utils.RandString(32)
	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	conn.Do("set", constant.USER_FLAG+token, userJson)
	utils.Expire(constant.USER_FLAG+token, 30)
	defer conn.Close()
	return dto.LoginResponse{StatusCode: 0, StatusMsg: "登录成功", UserId: user.UserId, Token: token}
}

//获取用户信息
func GetUserInfo(token string, userId int) dto.UserResponse {
	if token == "" || userId < 0 {
		return dto.UserResponse{StatusCode: -1, StatusMsg: "无效参数"}
	}
	conn := utils.GetRedisConnection()
	userjson, err := redis.String(conn.Do("get", constant.USER_FLAG+token))
	if err != nil {
		fmt.Println(err)
	}

	var user common.User
	json.Unmarshal([]byte(userjson), &user)
	//redis中不存在用户信息，从数据库中查找
	if user.Username == "" {
		user = dao.GetInfos(userId)
	}

	defer conn.Close()
	return dto.UserResponse{StatusCode: 0, StatusMsg: "操作成功", User: user}

}
