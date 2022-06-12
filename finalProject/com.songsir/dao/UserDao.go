package dao

import (
	"finalProject/com.songsir/common"
	"fmt"
)

func GetUsers() {
	var users []common.User
	connect := getConnect()
	if connect == nil {
		fmt.Println("数据库连接失败")
		return
	}

	connect.Find(&users)

	fmt.Println(users[0].UserId)
}

//用户注册
func Register(username string, password string, userId int64) int {
	var users []common.User
	var result = 0
	connect := getConnect()
	connect.Where("username=?", username).Find(&users)

	if len(users) > 0 {
		//该用户名有人注册
		return result
	}

	var user = common.User{UserId: int(userId), Username: username, Password: password}
	connect.Create(&user)
	result = 1
	return result
}

//登录
func Login(username string) common.User {
	var users []common.User
	connect := getConnect()
	connect.Where("username=?", username).Find(&users)

	if len(users) > 0 {
		return users[0]
	}
	return common.User{}

}

//返回用户数据

func GetInfos(userId int) common.User {
	connect := getConnect()
	var user common.User
	connect.Where("user_id=?", userId).Find(&user)
	return user
}
