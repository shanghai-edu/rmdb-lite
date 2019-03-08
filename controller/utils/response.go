package utils

import (
	"github.com/shanghai-edu/rmdb-lite/g"
)

type ApiResult struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Version string      `json:"version"`
	Data    interface{} `json:"data"`
}

const (
	OK                  = 0
	InvalidAPIKEY       = 10001
	BodyJsonDecodeError = 20001
	RecordNotFound      = 30001
	RecordAlreadyExists = 30002
	InternalAPIError    = 50001
)

var CodeMsg = map[int64]string{
	OK:                  "ok",
	BodyJsonDecodeError: "请求体 JSON 编码错误",
	InvalidAPIKEY:       "不正确的 API KEY",
	RecordNotFound:      "找不到对应的记录",
	RecordAlreadyExists: "记录已经存在",
	InternalAPIError:    "服务器内部错误",
}

func ErrorRes(code int64) (res ApiResult) {
	res.Code = code
	res.Message = CodeMsg[code]
	res.Version = g.VERSION
	return
}

func SuccessRes(data interface{}) (apiResult ApiResult) {
	apiResult.Code = 0
	apiResult.Message = "ok"
	apiResult.Version = g.VERSION
	apiResult.Data = data
	return
}
