package service

import (
	"master-gin/dao/mysql"
	"master-gin/models"
	"master-gin/pkg/jwt"
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

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// NOTE: pass pointer
	userEntity, err := mysql.Login(user)
	if err != nil {
		return "", err
	}
	return jwt.GenToken(userEntity.UserID, userEntity.Username)
}
