package models

type ParamSigUup struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	RepPassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	// NOTE: get user_id from request
	//UserID
	PostID string `json:"post_id" binding:"required"`
	// NOTE: 1 : approval, 0: cancel, -1: against
	Direction int8 `json:"direction,string" binding:"oneof=1 0 -1"`
}

// NOTE: get query string
type ParamPostList struct {
	CommunityID int64  `form:"community_id"`
	Page        int64  `form:"page"`
	PageSize    int64  `form:"size"`
	Order       string `form:"order"`
}
