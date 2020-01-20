package result

import (
	"us-stock-trade-date/result/code"
)

// 响应体基础序列化器
type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (response Response) Result() Response {
	if response.Code == 0 {
		response.Code = code.Success
	}
	return response
}

// 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"trackId"`
}
