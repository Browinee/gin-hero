package service

import (
	"master-gin/dao/mysql"
	"master-gin/models"
	"master-gin/pkg/snowflake"
)

func SignUp(p *models.ParamSigUup) {
	mysql.QueryUserByUsername()
	snowflake.GenID()
	mysql.InsertUser()

}
