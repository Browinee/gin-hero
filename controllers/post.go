package controllers

import (
	"master-gin/models"
	"master-gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {

	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON error", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseError(c, http.StatusOK, CodeInvalidParam)
		return
	}

	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, http.StatusOK, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	if err := service.CreatePost(p); err != nil {
		zap.L().Error("service.CreatePost failed", zap.Error((err)))
		ResponseError(c, http.StatusOK, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

func GetPostDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, http.StatusOK, CodeInvalidParam)
		return
	}
	data, err := service.GetPostById(id)
	if err != nil {
		zap.L().Error("service.GetPostDetailHandler error", zap.Error(err))
		ResponseError(c, http.StatusOK, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}
