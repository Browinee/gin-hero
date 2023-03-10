package controllers

import (
	"master-gin/models"
	"master-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type VoteData struct {
	// NOTE:get UserID from c.Get(controllers.ContextUserID)
	// UserID int64
	PostID int64 `json:"post_id,string" binding:"required"`
	// NOTE: 1: approve, -1 not approve, 0: cancel approve
	Direction int8 `json:"direction,string" binding:"required,oneof=1 0 -1"`
}

func PostVoteController(c *gin.Context) {
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		// NOTE: reference https://tw511.com/a/01/14663.html
		// There are three kinds of errors
		// ValidationErrors,InvalidValidationError, nil
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, http.StatusOK, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, http.StatusOK, CodeInvalidParam, errData)
		return
	}
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, http.StatusUnauthorized, CodeNeedLogin)
		return
	}
	if err := service.VoteForPost(userID, p); err != nil {
		zap.L().Error("service.VoteForPost", zap.Error(err))
		ResponseError(c, http.StatusOK, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
