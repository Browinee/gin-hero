package mysql

import (
	"database/sql"
	"master-gin/models"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	err = db.Select(&communityList, sqlStr)
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
