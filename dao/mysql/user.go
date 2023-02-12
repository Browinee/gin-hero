package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"master-gin/models"

	"go.uber.org/zap"
)

var (
	ErrorUserExist                   = errors.New("user existed")
	ErrorIncorrectUsernameOrPassword = errors.New("Incorrect username or password")
)

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username= ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	fmt.Println("count", count)
	if count > 0 {
		return ErrorUserExist
	}
	return
}
func QueryUserByUsername() {}

func InsertUser(user *models.User) (err error) {
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	user.Password = encryptPassword(user.Password)
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

const secret = "12345"

func encryptPassword(oPwd string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPwd)))
}

func Login(user *models.User) (userEntity *models.User, err error) {

	sqlStr := `select user_id, username, password from user where username=?`
	userFromDB := new(models.User)
	err = db.Get(userFromDB, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		zap.L().Error("user is not existed. ", zap.Error(err))
		return nil, ErrorIncorrectUsernameOrPassword

	}
	if err != nil {
		// NOTE: database got err when searching
		return nil, err
	}
	passwordFromUser := user.Password
	password := encryptPassword(passwordFromUser)
	if password != userFromDB.Password {
		zap.L().Error("incorrect password", zap.Error(err))
		return nil, ErrorIncorrectUsernameOrPassword
	}
	return userFromDB, nil
}
