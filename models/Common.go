package models

type JsonResultCode int

const (
	JRCodeSucc JsonResultCode = iota
	JRCodeFailed
	JRCode302 = 302 //跳转至地址
	JRCode401 = 401 //未授权访问
)

const (
	Deleted = iota - 1
	Disabled
	Enabled
)

// JsonResult 用于返回ajax请求的基类
type JsonResult struct {
	Code JsonResultCode `json:"code"`
	Msg  string         `json:"msg"`
	Data interface{}    `json:"data"`
}

type ListJsonResult struct {
	Code  JsonResultCode `json:"code"`
	Msg   string         `json:"msg"`
	Count int64          `json:"count"`
	Data  interface{}    `json:"data"`
}
