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
	"strconv"
)

func VideoFavoriteAction(token string, videoId string, actionType string) dto.UploadResponse {
	video, err := strconv.Atoi(videoId)
	if err != nil {
		fmt.Println(err)
	}
	action, err := strconv.Atoi(actionType)
	if err != nil {
		fmt.Println(err)
	}
	if video < 0 || action < 0 || action > 2 || len(token) == 0 {
		return dto.UploadResponse{StatusCode: -1, StatusMsg: "无效参数"}
	}
	//redis连接
	conn := utils.GetRedisConnection()
	defer conn.Close()
	//获取用户
	var user common.User
	jsonstr, err := redis.String(conn.Do("get", constant.USER_FLAG+token))
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal([]byte(jsonstr), &user)
	//点赞
	if action == 1 {
		//redis中set判断是否存在
		result, err := redis.Int(conn.Do("SISMEMBER", constant.VIDEO_FLAG+string(videoId), user.UserId))
		//不存在，点赞操作
		if err != nil {
			fmt.Println(err)
		}
		if result == 0 {
			conn.Do("SADD", constant.VIDEO_FLAG+string(videoId), user.UserId)
			dao.CreateFollow(video, user.UserId)
			dao.ActionNum("favorite", 1, videoId)
		}
	} else {
		//存在，取消点赞
		conn.Do("sRem", constant.VIDEO_FLAG+string(videoId), user.UserId)
		//删除数据库点赞信息
		dao.DeleteVideoFollow(video, user.UserId)
		dao.ActionNum("favorite", -1, videoId)
	}
	//加锁
	utils.SetByUser("setnx", constant.VIDEO_UPDATE_KEY, constant.VIDEO_UPDATE_VALUE)

	return dto.UploadResponse{StatusCode: 0, StatusMsg: "操作成功"}
}

func VideoFavoriteActionList(token string, user string) dto.PublishListResponse {

	if len(token) < 0 || len(user) < 0 {
		return dto.PublishListResponse{StatusCode: -1, StatusMsg: "参数无效"}
	}

	userId, err := strconv.Atoi(user)
	if err != nil {
		fmt.Println(err)
	}

	//查看点赞列表
	videoFollowList := dao.GetVideoFollowList(userId)
	//查询视频信息
	var videoList []common.Video
	var i = 0
	for i = 0; i < len(videoFollowList); i++ {
		video := dao.SelectOne(userId)
		videoList = append(videoList, video)
	}

	return dto.PublishListResponse{StatusCode: 0, StatusMsg: "请求成功", VideoList: videoList}
}
