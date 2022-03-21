package msg

import "net/http"

/*
创建人员：云深不知处
创建时间：2022/3/14
程序功能：
*/

const (
	SuccessCode = 1
	ErrCode     = 0
)

// API Response 基础序列化器
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

// Err 通用错误处理
func ErrResp(errStr string) (int, Response) {
	res := Response{
		Code:  ErrCode,
		Data:  nil,
		Error: errStr,
	}
	return http.StatusOK, res
}

// SuccessResp 通用处理
func SuccessResp(data interface{}) (int, Response) {
	res := Response{
		Code:  SuccessCode,
		Data:  data,
		Error: "",
	}
	return http.StatusOK, res
}
