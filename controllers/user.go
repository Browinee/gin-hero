package controllers

import (
	"fmt"
	"master-gin/models"
	"master-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	// var p *models.ParamSigUp or the following
	p := new(models.ParamSigUup)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Signup with invalid param", zap.Error(err))

		// NOTE: check err is validator error
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, http.StatusUnauthorized, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, http.StatusUnauthorized, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	if err := service.SignUp(p); err != nil {
		zap.L().Error("service.signup failed", zap.Error(err))
		ResponseErrorWithMsg(c, http.StatusUnauthorized, CodeInvalidParam, "Register failed")
		return
	}

	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))

		// NOTE: check err is validator error
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return

	}
	fmt.Println("123")
	token, err := service.Login(p)
	if err != nil {
		zap.L().Error("service.login failed", zap.String("username", p.Username), zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "Incorrect username or password.",
		})
		return
	}
	fmt.Println("token", token)
	ResponseSuccess(c, token)
	return
}
