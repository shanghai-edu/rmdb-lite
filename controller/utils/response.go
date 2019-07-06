package utils

import (
	"github.com/shanghai-edu/rmdb-lite/g"
)

//APIResult 接口返回的标准结构
type APIResult struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Version string      `json:"version"`
	Data    interface{} `json:"data"`
}

const (
	//OK 状态正常的返回码
	OK = 0
	//InvalidAPIKEY X-API-KEY 非法的返回码
	InvalidAPIKEY = 10001
	//BodyJSONDecodeError JSON 解码失败的返回码
	BodyJSONDecodeError = 20001
	//RecordNotFound 没有查到数据的返回码
	RecordNotFound = 30001
	//RecordAlreadyExists 插入数据时，数据已经存在的返回码
	RecordAlreadyExists = 30002
	//InternalAPIError 内部服务错误
	InternalAPIError = 50001
)

//CodeMsg 返回码对应的中文 MSG 说明
var CodeMsg = map[int64]string{
	OK:                  "ok",
	BodyJSONDecodeError: "请求体 JSON 编码错误",
	InvalidAPIKEY:       "不正确的 API KEY",
	RecordNotFound:      "找不到对应的记录",
	RecordAlreadyExists: "记录已经存在",
	InternalAPIError:    "服务器内部错误",
}

//ErrorRes 错误返回
func ErrorRes(code int64) (res APIResult) {
	res.Code = code
	res.Message = CodeMsg[code]
	res.Version = g.VERSION
	return
}

//SuccessRes 正确返回
func SuccessRes(data interface{}) (apiResult APIResult) {
	apiResult.Code = 0
	apiResult.Message = "ok"
	apiResult.Version = g.VERSION
	apiResult.Data = data
	return
}
