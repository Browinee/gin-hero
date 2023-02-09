package service

import (
	"master-gin/dao/mysql"
	"master-gin/models"
	"master-gin/pkg/snowflake"
)

func SignUp(p *models.ParamSigUup) (err error) {

	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	userID := snowflake.GenID()
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.InsertUser(user)

}
