package controllers

import (
	"errors"
	"master-gin/middlewares"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("User not login.")

func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(middlewares.ContextUserID)
	if !ok {
		err = ErrorUserNotLogin
		return
	}

	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
