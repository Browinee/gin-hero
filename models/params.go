package models

type ParamSigUup struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	RepPassword string `json:"re_password" binding:"required,eqfield=Password"`
}
