package controllers

import (
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
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Incorrect username or password.",
			})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})

		return
	}

	service.SignUp(p)
	c.JSON(http.StatusOK, gin.H{
		"msg": "success.",
	})
}
