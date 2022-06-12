package dto

import "finalProject/com.songsir/common"

type LoginDto struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type LoginResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int    `json:"user_id"`
	Token      string `json:"token"`
}

type UserResponse struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	User       common.User `json:"user"`
}

type PublishListResponse struct {
	StatusCode int            `json:"status_code"`
	StatusMsg  string         `json:"status_msg"`
	NextTime   int            `json:"next_time"`
	VideoList  []common.Video `json:"video_list"`
}

type UploadResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type CommentInfoResponse struct {
	StatusCode int                   `json:"status_code"`
	StatusMsg  string                `json:"status_msg"`
	Comment    []common.VideoComment `json:"comment"`
}

type CommentListResponse struct {
	StatusCode int                   `json:"status_code"`
	StatusMsg  string                `json:"status_msg"`
	Comment    []common.VideoComment `json:"comment_list"`
}
