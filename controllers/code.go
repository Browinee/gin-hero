package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeInvalidToken
	CodeNeedLogin
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "Success",
	CodeInvalidParam:    "Invalid param",
	CodeUserExist:       "User existed",
	CodeUserNotExist:    "User not existed",
	CodeInvalidPassword: "Invalid Password",
	CodeServerBusy:      "Server busy",
	CodeInvalidToken:    "Invalid token",
	CodeNeedLogin:       "Please login",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
