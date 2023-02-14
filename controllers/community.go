package controllers

import (
	"master-gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	data, err := service.GetCommunityList()
	if err != nil {
		zap.L().Error("service.GetCommunityList", zap.Error(err))
		ResponseError(c, http.StatusOK, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		ResponseError(c, http.StatusOK, CodeInvalidParam)
		return
	}

	data, err := service.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("service.")
		ResponseError(c, http.StatusOK, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}
