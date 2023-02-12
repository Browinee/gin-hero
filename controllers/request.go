package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const ContextUserID = "userID"

var ErrorUserNotLogin = errors.New("User not login.")

func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(ContextUserID)
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
