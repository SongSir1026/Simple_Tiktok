package dao

import (
	"finalProject/com.songsir/common"
	"time"
)

func CreateFollow(videoId int, userId int) {

	id := time.Now().Unix()

	var videoFollow = common.VideoFollow{Id: int(id), VideoId: videoId, UserId: userId}

	con := getConnect()
	con.Table("video_follow").Create(&videoFollow)

}

func DeleteVideoFollow(videoId int, userId int) {
	connect := getConnect()
	var videoFollow common.VideoFollow
	connect.Table("video_follow").Where("video_id=?", videoId).Where("user_id=?", userId).Find(&videoFollow)
	connect.Table("video_follow").Delete(&videoFollow)
}

func SelectFollowOne(videoId int, userId int) common.VideoFollow {
	connect := getConnect()
	var videoFollow common.VideoFollow
	connect.Table("video_follow").Where("video_id=?", videoId).Where("user_id=?", userId).Find(&videoFollow)
	return videoFollow
}

func ActionNum(sql string, action int, videoId string) {
	conn := getConnect()

	var video common.Video

	conn.Where("video_id", videoId).Find(&video)

	if sql == "favorite" {
		video.FavoriteCount += action

		conn.Model(&video).Where("video_id", videoId).Update("favorite_count", video.FavoriteCount)

	} else if sql == "comment" {
		video.CommentCount += action
		conn.Model(&video).Where("video_id", videoId).Update("comment_count", video.CommentCount)
	}
}

func GetVideoFollowList(userId int) []common.VideoFollow {
	conn := getConnect()
	var videolist []common.VideoFollow
	conn.Where("user_id", userId).Find(&videolist)
	return videolist
}
