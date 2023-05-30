package tools

import "time"

const (
	OK                  = 0
	UserInfoError       = 10001
	CaptchaError        = 10002
	UserExist           = 10003
	UserIsNotExist      = 10004
	SourceExist         = 10005
	NoLogin             = 10006
	BadRequest          = 10400
	InternalServerError = 10500
	// 借书时长30天
	T = 30 * 86400 * time.Second
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
