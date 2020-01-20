package code

// 定义所有状态码
const (
	// 成功
	Success = 0

	// 参数缺失
	ParamMiss = 100001

	// 参数无效
	ParamInvalid = 100002

	// JSON 解析出错
	JsonUnmarshalError = 100003

	// 年份超出范围
	YearOutOfRange = 200001

	// 系统错误
	SystemError = 999999
)
