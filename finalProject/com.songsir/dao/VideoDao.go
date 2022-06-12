package dao

import (
	"finalProject/com.songsir/common"
	"finalProject/com.songsir/common/constant"
	"finalProject/com.songsir/common/utils"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

//查询投稿列表
func SelectPublishList(userId int) []common.Video {
	connect := getConnect()
	var videolist []common.Video
	connect.Where("author_id", userId).Find(&videolist)
	return videolist
}

//获取视频流
func SelectVideo() []common.Video {
	connect := getConnect()
	var videolist []common.Video
	connect.Find(&videolist)
	conn := utils.GetRedisConnection()
	//redis中set判断是否存在

	var i = 0
	for i = 0; i < len(videolist); i++ {

		user := GetInfos(videolist[i].AuthorId)
		videolist[i].User = user

		result := 1
		result, err := redis.Int(conn.Do("SISMEMBER", constant.VIDEO_FLAG+string(videolist[i].VideoId), user.UserId))
		fmt.Println(result)
		//不存在，点赞操作
		if err != nil {
			fmt.Println(err)
		}
		one := SelectFollowOne(videolist[i].VideoId, user.UserId)

		if result == 0 && one.VideoId > 0 {
			videolist[i].IsFavorite = true
		}
	}
	return videolist

}

//上传视频
func InsertVideo(title string, playUrl string, userId int) {

	unix := time.Now().Unix()
	var coverUrl = playUrl + "?x-oss-process=video/snapshot,t_500,f_jpg,w_600,h_800,m_fast"
	var video = common.Video{VideoId: int(unix), Title: title, AuthorId: userId, PlayUrl: playUrl, CoverUrl: coverUrl}

	conn := getConnect()

	conn.Create(&video)
}

//查询一个
func SelectOne(videoId int) common.Video {
	connect := getConnect()
	var video common.Video
	connect.Where("video_id", videoId).Find(&video)
	user := GetInfos(video.AuthorId)
	video.User = user
	return video
}
