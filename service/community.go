package service

import (
	"master-gin/dao/mysql"
	"master-gin/models"
)

func GetCommunicityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}
