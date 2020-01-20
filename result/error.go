package result

import (
	"encoding/json"
	"fmt"
	"us-stock-trade-date/config"
	"us-stock-trade-date/result/code"

	"gopkg.in/go-playground/validator.v8"
)

// 解析错误转为响应体
func ErrorResponse(err error) Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := config.T(fmt.Sprintf("Field.%s", e.Field))
			tag := config.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return Response{
				Code:    code.ParamInvalid,
				Message: fmt.Sprintf("%s%s", field, tag),
			}
		}
	}

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return Response{
			Code:    code.JsonUnmarshalError,
			Message: err.(*json.UnmarshalTypeError).Field + "参数的类型出错",
		}
	}

	if _, ok := err.(*json.SyntaxError); ok {
		return Response{
			Code:    code.JsonUnmarshalError,
			Message: "JSON 无法被解析",
		}
	}

	return Response{
		Code:    code.SystemError,
		Message: "系统错误",
	}
}
