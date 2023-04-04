package utils

import (
	"encoding/json"
	"noir.github.com/aws_billing/internal/modules/logger"
)

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

type JsonResponse struct{}

const ResponseSuccess = 200
const ResponseFailure = 400
const NotFound = 404
const ServerError = 500

const SuccessContent = "操作成功"
const FailureContent = "操作失败"

func (j *JsonResponse) SuccessJson(msg string, data interface{}) map[string]interface{} {
	return j.responseJson(ResponseSuccess, msg, data)
}

func (j *JsonResponse) FailureJson(code int, msg string) map[string]interface{} {
	return j.responseJson(code, msg, nil)
}

func (j *JsonResponse) CommonFailureJson(msg string, err ...error) map[string]interface{} {
	if len(err) > 0 {
		logger.Warn(err)
	}
	return j.FailureJson(ResponseFailure, msg)
}

func (j *JsonResponse) responseJson(code int, msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    code,
		"message": msg,
		"data":    data,
	}
}

func JsonResponseByErr(err error) string {
	jsonResp := JsonResponse{}
	if err != nil {
		return jsonResp.CommonFailure(FailureContent, err)
	}

	return jsonResp.Success(SuccessContent, nil)
}

func (j *JsonResponse) Success(msg string, data interface{}) string {
	return j.response(ResponseSuccess, msg, data)
}

func (j *JsonResponse) Failure(code int, msg string) string {
	return j.response(code, msg, nil)
}

func (j *JsonResponse) CommonFailure(msg string, err ...error) string {
	if len(err) > 0 {
		logger.Warn(err)
	}
	return j.Failure(ResponseFailure, msg)
}

func (j *JsonResponse) response(code int, msg string, data interface{}) string {
	resp := response{
		code,
		msg,
		data,
	}
	result, err := json.Marshal(resp)
	if err != nil {
		logger.Error(err)
	}
	return string(result)
}
