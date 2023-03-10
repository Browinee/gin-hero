package controllers

import (
	"master-gin/constants"
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
		ResponseError(c, http.StatusBadRequest, CodeInvalidParam)
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
		ResponseError(c, http.StatusBadRequest, CodeInvalidParam)
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

func GetPostListHandler(c *gin.Context) {
	offset, limit := getPageInfo(c)
	data, err := service.GetPostList(offset, limit)
	if err != nil {
		zap.L().Error("service.GetPostListHandler error", zap.Error(err))
		ResponseError(c, http.StatusOK, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

// NOTE: sorted by create_time and vote
func GetPostListHandler2(c *gin.Context) {
	p := &models.ParamPostList{
		Page:     1,
		PageSize: 10,
		Order:    constants.DefaultOrderTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(c, http.StatusBadRequest, CodeInvalidParam)
		return
	}

	data, err := service.GetPostListNew(p)
	if err != nil {
		zap.L().Error("service.GetPostListHandler2 error", zap.Error(err))
		ResponseError(c, http.StatusOK, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// func GetCommunityPostListHandler(c *gin.Context) {
// 	p := &models.ParamPostList{
// 		Page:        1,
// 		PageSize:    10,
// 		Order:       constants.DefaultOrderTime,
// 		CommunityID: 0,
// 	}
// 	if err := c.ShouldBindQuery(p); err != nil {
// 		zap.L().Error("GetCommunityPostListHandler with invalid params", zap.Error(err))
// 		ResponseError(c, http.StatusBadRequest, CodeInvalidParam)
// 		return
// 	}
// 	data, err := service.GetCommunityPostList(p)
// 	if err != nil {
// 		ResponseError(c, http.StatusOK, CodeServerBusy)
// 		return
// 	}
// 	ResponseSuccess(c, data)
// }
