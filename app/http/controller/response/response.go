package response

import (
	"context"
)

var (
	SUCCESS           = resData{0, "成功"}
	INVALID_PARAMETER = resData{1001, "无效的参数"}
	INVALID_IDENTITY  = resData{1002, "身份验证失败"}
	PERMISSION_DENIED = resData{1003, "没有权限"}
	RESULT_EMPTY      = resData{1004, "没有结果"}
	SERVER_ERROR      = resData{3001, "服务器错误"}
	TIMEOUT           = resData{4001, "访问超时"}
)

type resData struct {
	no  int
	msg string
}
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Success(ctx context.Context, data ...interface{}) Response {
	return getResponse(ctx, SUCCESS, "", data...)
}
func SuccessWithMsg(ctx context.Context, msg string, data ...interface{}) Response {
	return getResponse(ctx, SUCCESS, msg, data...)
}
func Failure(ctx context.Context, resErr resData, msg string, data ...interface{}) Response {
	return getResponse(ctx, resErr, msg, data...)
}
func getResponse(ctx context.Context, resErr resData, msg string, data ...interface{}) Response {
	resMsg := resErr.msg

	if msg != "" {
		resMsg += "," + msg
	}
	if len(data) > 0 {
		return Response{
			Code: resErr.no,
			Msg:  resMsg,
			Data: data[0],
		}
	}
	return Response{
		Code: resErr.no,
		Msg:  resMsg,
	}
}
