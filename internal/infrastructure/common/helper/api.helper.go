package helper

import (
	"fmt"
	"net/http"
	"reflect"

	cons "restuwahyu13/shopping-cart/internal/domain/constant"
	inf "restuwahyu13/shopping-cart/internal/domain/interface/helper"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
)

var errorCodeMapping = map[int]string{
	http.StatusBadGateway:          "SERVICE_ERROR",
	http.StatusServiceUnavailable:  "SERVICE_UNAVAILABLE",
	http.StatusGatewayTimeout:      "SERVICE_TIMEOUT",
	http.StatusConflict:            "DUPLICATE_RESOURCE",
	http.StatusBadRequest:          "INVALID_REQUEST",
	http.StatusUnprocessableEntity: "INVALID_REQUEST",
	http.StatusPreconditionFailed:  "REQUEST_COULD_NOT_BE_PROCESSED",
	http.StatusForbidden:           "ACCESS_DENIED",
	http.StatusUnauthorized:        "UNAUTHORIZED_TOKEN",
	http.StatusNotFound:            "UNKNOWN_RESOURCE",
	http.StatusInternalServerError: "GENERAL_ERROR",
}

func Version(path string) string {
	return fmt.Sprintf("%s/%s", cons.API, path)
}

func Api(rw http.ResponseWriter, options hopt.Response) {
	response := buildResponse(options)
	writeResponse(rw, NewParser(), response)
}

func getErrorCode(statusCode int) string {
	if code, exists := errorCodeMapping[statusCode]; exists {
		return code
	}

	return errorCodeMapping[http.StatusInternalServerError]
}

func isEmptyResponse(resp hopt.Response) bool {
	return reflect.DeepEqual(resp, hopt.Response{})
}

func buildResponse(options hopt.Response) hopt.Response {
	response := hopt.Response{
		StatCode: http.StatusInternalServerError,
		ErrMsg:   cons.DEFAULT_ERR_MSG,
	}

	if isEmptyResponse(options) {
		defaultErrCode := getErrorCode(http.StatusInternalServerError)
		response.ErrCode = &defaultErrCode

		return response
	}

	response = copyResponseFields(options, response)
	setResponseDefaults(&response)

	return response
}

func copyResponseFields(source, target hopt.Response) hopt.Response {
	if source.StatCode != 0 {
		target.StatCode = source.StatCode
	}

	if source.Message != nil {
		target.Message = source.Message
	}

	if source.ErrCode != nil {
		target.ErrCode = source.ErrCode
	}

	if source.ErrMsg != "" {
		target.ErrMsg = source.ErrMsg
	}

	if source.Data != nil {
		target.Data = source.Data
	}

	if source.Errors != nil {
		target.Errors = source.Errors
	}

	if source.Pagination != nil {
		target.Pagination = source.Pagination
	}

	target = hopt.Response{
		StatCode:   target.StatCode,
		Message:    target.Message,
		ErrCode:    target.ErrCode,
		ErrMsg:     target.ErrMsg,
		Data:       target.Data,
		Errors:     target.Errors,
		Pagination: target.Pagination,
	}

	return target
}

func setResponseDefaults(response *hopt.Response) {
	if response.StatCode == 0 {
		response.StatCode = http.StatusInternalServerError
	}

	if response.StatCode >= http.StatusBadRequest && response.ErrCode == nil {
		defaultErrCode := getErrorCode(int(response.StatCode))
		response.ErrCode = &defaultErrCode
	}

	if response.StatCode >= http.StatusInternalServerError && response.ErrMsg == cons.DEFAULT_ERR_MSG {
		response.ErrMsg = cons.DEFAULT_ERR_MSG
	}
}

func writeResponse(rw http.ResponseWriter, parser inf.IParser, response hopt.Response) {
	rw.Header().Set("Content-Type", "application/json")

	statusCode := response.StatCode
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	rw.WriteHeader(int(statusCode))

	if err := parser.Encode(rw, response); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)

		errorResponse := fmt.Sprintf(`{"stat_code":%d, "err_code":"%s", "err_msg":"%s"}`, http.StatusInternalServerError, errorCodeMapping[http.StatusInternalServerError], cons.DEFAULT_ERR_MSG)
		fmt.Fprint(rw, errorResponse)
	}
}
