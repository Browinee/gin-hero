package service

import (
	"master-gin/dao/mysql"
	"master-gin/models"
	"master-gin/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	p.ID = snowflake.GenID()

	return mysql.CreatePost(p)
}
