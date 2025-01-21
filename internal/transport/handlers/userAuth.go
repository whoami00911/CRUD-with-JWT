package handlers

import (
	"fmt"
	"webPractice1/internal/domain"

	"github.com/gin-gonic/gin"
)

func (har *HandlerAssetsResponse) singUp(c *gin.Context) {
	var input domain.User
	if err := c.BindJSON(&input); err != nil {
		har.Logger.Error(fmt.Sprintf("BindJSON method error: %s", err))
		return
	}
}

func (har *HandlerAssetsResponse) singIn(c *gin.Context) {

}
