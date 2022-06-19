package service

import (
	"finalProject/com.songsir/common"
	"finalProject/com.songsir/common/constant"
	"finalProject/com.songsir/common/utils"
	"finalProject/com.songsir/dao"
	"finalProject/com.songsir/dto"
	"fmt"
	"regexp"
	"strconv"
)

func CommentAction(token string, video string, action string, commenttext string, comm string) dto.CommentInfoResponse {

	//防止Sql注入
	m1, err1 := regexp.MatchString(constant.SQL_JUDGE, commenttext)
	if err1 != nil {
		fmt.Println(err1)
	}
	if m1 {
		return dto.CommentInfoResponse{StatusMsg: "无效账号或者密码", StatusCode: -1}
	}
	videoId, err := strconv.Atoi(video)
	if err != nil {
		fmt.Println(err)
	}

	actionType, err := strconv.Atoi(action)
	if err != nil {
		fmt.Println(err)
	}

	//判断类别
	if actionType >= 2 || actionType < 0 {
		commentId, err := strconv.Atoi(comm)
		if err != nil {
			fmt.Println(err)
		}
		//删除评论
		dao.DeleteComment(videoId, commentId)
		dao.ActionNum(constant.COMMENT_FLAG, -1, video)
		return dto.CommentInfoResponse{StatusCode: 0, StatusMsg: "删除成功"}
	}
	dao.AddComment(videoId, commenttext)
	var commentList []common.VideoComment
	commentList = dao.ShowComment(videoId)
	dao.ActionNum(constant.COMMENT_FLAG, 1, video)
	utils.SetByUser("setnx", constant.VIDEO_UPDATE_KEY, constant.VIDEO_UPDATE_VALUE)
	return dto.CommentInfoResponse{StatusCode: 0, StatusMsg: "请求成功", Comment: commentList}
}

func GetCommentList(token string, video string) dto.CommentListResponse {
	videoId, err := strconv.Atoi(video)
	if err != nil {
		fmt.Println(err)
	}

	commentList := dao.ShowComment(videoId)

	return dto.CommentListResponse{StatusCode: 0, StatusMsg: "请求成功", Comment: commentList}

}
