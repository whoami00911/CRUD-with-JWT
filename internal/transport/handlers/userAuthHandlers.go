package handlers

import (
	"errors"
	"fmt"
	"webPractice1/internal/domain"
	"webPractice1/internal/service"

	"github.com/gin-gonic/gin"
)

func (har *HandlerAssetsResponse) singUp(c *gin.Context) {
	var input domain.User
	if err := c.BindJSON(&input); err != nil {
		har.Logger.Error(fmt.Sprintf("BindJSON method error: %s", err))
		c.JSON(400, gin.H{"error": "Invalid data input"})
		return
	}
	id := har.service.Autherization.CreateUser(input)
	c.JSON(200, map[string]int{
		"id": id,
	})
}

func (har *HandlerAssetsResponse) singIn(c *gin.Context) {
	var input *domain.UserSignIn
	if err := c.BindJSON(&input); err != nil {
		har.Logger.Error(fmt.Sprintf("BindJSON method error: %s", err))
		c.JSON(400, gin.H{"error": "Invalid data input"})
		return
	}
	token, err := har.service.Autherization.GenToken(input.Username, input.Password)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			c.JSON(400, gin.H{"error": "user or password is incorrect"})
			return
		}
		c.JSON(500, gin.H{"error": "something wrong"})
		return
	}
	//c.Header("Authorization", "Bearer "+token) //отправить в хедере
	c.JSON(200, map[string]string{
		"token": token,
	})
}
