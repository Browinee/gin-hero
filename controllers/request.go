package controllers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const ContextUserID = "userID"

var ErrorUserNotLogin = errors.New("User not login.")

func getCurrentUserID(c *gin.Context) (userID int64, err error) {
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

func getPageInfo(c *gin.Context) (offset, limit int64) {
	offsetStr := c.Query("page")
	limitStr := c.Query("pageSize")

	var (
		err error
	)

	offset, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 1
	}
	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 10
	}
	return offset, limit
}
