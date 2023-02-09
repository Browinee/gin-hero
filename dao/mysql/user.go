package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"master-gin/models"
)

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username= ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	fmt.Println("count", count)
	if count > 0 {
		return errors.New("user existed")
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
