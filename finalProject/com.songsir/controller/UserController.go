package controller

import (
	"finalProject/com.songsir/form"
	"finalProject/com.songsir/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var loginForm form.LoginForm
	c.BindJSON(&loginForm)
	result := service.Register(loginForm.Username, loginForm.Password)
	fmt.Println("insert person username {}", loginForm.Username)
	c.JSON(http.StatusOK, gin.H{"msg": result})
}
