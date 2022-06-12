package main

import (
	"finalProject/com.songsir/dao"
	_ "finalProject/com.songsir/dao"
	_ "finalProject/com.songsir/form"
	"finalProject/com.songsir/service"
	_ "finalProject/com.songsir/service"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

func main() {
	dao.GetUsers()
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works")
	})

	//注册接口
	router.POST("/douyin/user/register/", func(c *gin.Context) {
		//var loginForm form.LoginForm
		//c.BindJSON(&loginForm)
		username := c.Query("username")
		password := c.Query("password")

		result := service.Register(username, password)
		//result := service.Register(loginForm.Username, loginForm.Password)
		c.JSON(http.StatusOK, gin.H{"msg": result})
	})

	//登录接口
	router.POST("/douyin/user/login/", func(c *gin.Context) {
		//var loginForm form.LoginForm
		//c.BindJSON(&loginForm)
		var username = c.Query("username")
		var password = c.Query("password")
		result := service.Login(username, password)
		//result := service.Login(loginForm.Username, loginForm.Password)
		//fmt.Println("insert person username {}", loginForm.Username)
		c.JSON(http.StatusOK, result)
	})

	//获得用户信息
	router.GET("/douyin/user/", func(c *gin.Context) {
		userId := c.Query("user_id")
		token := c.Query("token")
		atoi, err := strconv.Atoi(userId)
		if err != nil {
			fmt.Println(err)
		}
		result := service.GetUserInfo(token, atoi)
		c.JSON(http.StatusOK, result)
	})

	//发布列表
	router.GET("/douyin/publish/list/", func(c *gin.Context) {
		userId := c.Query("user_id")
		token := c.Query("token")
		atoi, err := strconv.Atoi(userId)
		if err != nil {
			fmt.Println(err)
		}
		result := service.ShowPublishList(token, atoi)
		c.JSON(http.StatusOK, result)
	})

	//视屏频接口
	router.GET("/douyin/feed", func(c *gin.Context) {
		latestTime := c.Query("latest_time")
		token := c.Query("token")
		result := service.GetFeedStream(latestTime, token)

		c.JSON(http.StatusOK, result)
	})

	//视频投稿
	router.POST("/douyin/publish/action/", func(c *gin.Context) {
		title := c.PostForm("title")
		file, err := c.FormFile("data")
		token := c.PostForm("token")
		if err != nil {
			panic(err)
		}

		result := service.UploadVideo(title, file, token)
		c.JSON(http.StatusOK, result)
	})

	//点赞
	router.POST("/douyin/favorite/action/", func(c *gin.Context) {
		token := c.Query("token")
		videoId := c.Query("video_id")
		actionType := c.Query("action_type")
		result := service.VideoFavoriteAction(token, videoId, actionType)
		c.JSON(http.StatusOK, result)

	})

	//点赞列表
	router.GET("/douyin/favorite/list/", func(c *gin.Context) {
		token := c.Query("token")
		userId := c.Query("user_id")

		result := service.VideoFavoriteActionList(token, userId)
		c.JSON(http.StatusOK, result)
	})

	//评论视频
	router.POST("/douyin/comment/action/", func(c *gin.Context) {
		token := c.Query("token")
		videoId := c.Query("video_id")
		actionType := c.Query("action_type")
		commentText := c.Query("comment_text")
		commentId := c.Query("comment_id")
		result := service.CommentAction(token, videoId, actionType, commentText, commentId)
		c.JSON(http.StatusOK, result)
	})

	//评论列表
	router.GET("/douyin/comment/list/", func(c *gin.Context) {
		token := c.Query("token")
		videoId := c.Query("video_id")
		result := service.GetCommentList(token, videoId)
		c.JSON(http.StatusOK, result)
	})

	router.Run(":8080")
}
