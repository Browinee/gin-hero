package controllers

import (
	"master-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	data, err := service.GetCommunicityList()
	if err != nil {
		zap.L().Error("service.GetCommunicityList", zap.Error(err))
		ResponseError(c, http.StatusOK, CodeServerBusy)
	}
	ResponseSuccess(c, data)
}
