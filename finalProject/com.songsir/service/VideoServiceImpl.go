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
	"io/ioutil"
	"mime/multipart"
	"regexp"
)

func ShowPublishList(token string, userId int) dto.PublishListResponse {
	if token == "" || userId < 0 {
		return dto.PublishListResponse{StatusCode: -1, StatusMsg: "无效参数"}
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
	var videolist []common.Video
	videolist = dao.SelectPublishList(userId)

	//fmt.Println(videolist)
	return dto.PublishListResponse{StatusCode: 0, StatusMsg: "请求成功", VideoList: videolist}

}

func GetFeedStream(latestTime string, token string) dto.PublishListResponse {
	if token == "" || len(latestTime) == 0 {
		return dto.PublishListResponse{StatusCode: -1, StatusMsg: "无效参数"}
	}

	var videos []common.Video
	videos = dao.SelectVideo()
	//fmt.Println(videos)
	return dto.PublishListResponse{StatusCode: 0, StatusMsg: "请求成功", VideoList: videos}
}

func UploadVideo(title string, file *multipart.FileHeader, token string) dto.UploadResponse {

	//防止Sql注入
	m1, err1 := regexp.MatchString(constant.SQL_JUDGE, title)

	if err1 != nil {
		fmt.Println(err1)
	}
	if m1 {
		return dto.UploadResponse{StatusMsg: "无效账号或者密码", StatusCode: -1}
	}
	
	fileHandler, err := file.Open()
	if err != nil {
		return dto.UploadResponse{StatusCode: -1, StatusMsg: "文件打开错误"}
	}
	defer fileHandler.Close()

	fileByte, err := ioutil.ReadAll(fileHandler)

	if err != nil {
		return dto.UploadResponse{StatusCode: -1, StatusMsg: "文件上传失败"}
	}

	conn := utils.GetRedisConnection()

	str, err := redis.String(conn.Do("get", constant.USER_FLAG+token))
	if err != nil {
		fmt.Println(err)
	}

	var user common.User
	json.Unmarshal([]byte(str), &user)

	url, err := utils.UploadFile(file.Filename, fileByte, user.Username)

	if err != nil {
		return dto.UploadResponse{StatusCode: -1, StatusMsg: "文件上传失败"}
	}

	dao.InsertVideo(title, url, user.UserId)

	return dto.UploadResponse{StatusCode: 0, StatusMsg: "文件上传成功"}

}
