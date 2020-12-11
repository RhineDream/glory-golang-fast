package utils

/**
封装请求返回值
2020年12月11日
结构体中的声明变量首字母必须大写 不然无法被beego解析
*/
type ResponseResult struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"result"`
}

/**
成功返回,无参数
*/
func Ok() *ResponseResult {
	return &ResponseResult{200, "OK", nil}
}

/**
成功返回，只需要传数据
*/
func OkResult(data interface{}) *ResponseResult {
	return &ResponseResult{200, "OK", data}
}

/**
构建返回，可以成功，可以失败，需要传状态，提示信息
*/
func Build(status int, msg string) *ResponseResult {
	return &ResponseResult{status, msg, nil}
}

/**
构建返回，可以成功，可以失败，需要传状态，提示信息，数据
*/
func BuildResult(status int, msg string, data interface{}) *ResponseResult {
	return &ResponseResult{status, msg, data}
}

/**
构建失败，可以成功，可以失败，需要传状态，提示信息
*/
func Fail(status int, msg string) *ResponseResult {
	return &ResponseResult{status, msg, nil}
}

/**
构建失败，可以成功，可以失败，需要传状态，提示信息，数据
*/
func FailResult(status int, msg string, data interface{}) *ResponseResult {
	return &ResponseResult{status, msg, data}
}
