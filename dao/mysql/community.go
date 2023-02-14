package mysql

import (
	"database/sql"
	"master-gin/models"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			// NOTE: don't expose err to outside
			// use zap to log
			err = nil
		}
	}
	return
}

func GetCommunityDetailByID(id int64) (communityDetail *models.CommunityDetail, err error) {
	communityDetail = new(models.CommunityDetail)
	sqlStr := `select
					community_id, community_name, introduction, create_time
				from community
				 where community_id = ?
			`
	if err := db.Get(communityDetail, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no selected community in DB")
			err = ErrorInvalidID
		}
	}
	return communityDetail, err
}
