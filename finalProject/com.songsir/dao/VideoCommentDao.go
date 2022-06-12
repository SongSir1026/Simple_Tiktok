package dao

import (
	"finalProject/com.songsir/common"
	"time"
)

func AddComment(videoId int, commentText string) common.VideoComment {

	time := time.Now()
	id := time.Unix()

	var createDate = string(time.Format("01-02"))
	var videoComment = common.VideoComment{Id: int(id), VideoId: videoId, CommentText: commentText, CreateDate: createDate}

	conn := getConnect()
	conn.Create(&videoComment)
	return videoComment
}

func DeleteComment(videoId int, commentId int) {
	conn := getConnect()
	var videoComment common.VideoComment
	conn.Where("id=?", commentId).Where("video_id", videoId).Find(&videoComment)

	conn.Delete(&videoComment)
}

func ShowComment(videoId int) []common.VideoComment {
	var videoComment []common.VideoComment
	conn := getConnect()
	conn.Where("video_id", videoId).Find(&videoComment)
	var i = 0
	for i = 0; i < len(videoComment); i++ {
		user := GetInfos(videoComment[i].UserId)
		videoComment[i].User = user
	}
	return videoComment
}
