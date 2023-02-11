package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, status int, code ResCode) {
	responseData := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(status, responseData)
}

func ResponseErrorWithMsg(c *gin.Context, status int, code ResCode, msg interface{}) {
	responseData := &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, responseData)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	responseData := &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, responseData)
}
