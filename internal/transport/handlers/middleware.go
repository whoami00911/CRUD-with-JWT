package handlers

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func (har *HandlerAssetsResponse) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.JSON(401, "Invalid authorization token (is empty)")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.JSON(401, "Invalid authorization token")
		return
	}
	if len(headerParts[1]) == 0 {
		c.JSON(401, "Empty token")
		return
	}
	userId, err := har.service.ParseToken(headerParts[1])
	if err != nil {
		har.Logger.Error(fmt.Sprintf("failed to parse token: %s", err))
		return
	}
	c.Set("UserId", userId)
}

/*func (har *HandlerAssetsResponse) getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("Authorization")
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}*/
