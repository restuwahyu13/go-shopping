package helper

import (
	"fmt"
	"net/http"
	"reflect"

	cons "restuwahyu13/shopping-cart/internal/domain/constant"
	inf "restuwahyu13/shopping-cart/internal/domain/interface/helper"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
)

func Version(path string) string {
	return fmt.Sprintf("%s/%s", cons.API, path)
}

func Api(rw http.ResponseWriter, options hopt.Response) {
	var (
		parser     inf.IParser    = NewParser()
		res        hopt.Response  = hopt.Response{StatCode: http.StatusInternalServerError}
		config     map[string]any = nil
		errCode    string         = "GENERAL_ERROR"
		errMessage string         = cons.DEFAULT_ERR_MSG
	)

	if reflect.DeepEqual(options, hopt.Response{}) {
		res.ErrCode = errCode
		res.ErrMsg = errMessage
	}

	optionsByte, err := parser.Marshal(&options)
	if err != nil {
		res.ErrCode = errCode
		res.ErrCode = errMessage
	}

	if err := parser.Unmarshal(optionsByte, &config); err != nil {
		res.ErrCode = errCode
		res.ErrCode = errMessage
	}

	if (config["stat_code"] == nil && config["message"] == nil && config["err_msg"] != nil) || config["err_msg"] == nil || config["err_code"] == nil {
		res.ErrCode = errCode
		res.ErrMsg = config["err_msg"]
	}

	if statCode := config["stat_code"]; statCode != nil {
		res.StatCode = statCode.(float64)
	}

	for key, value := range config {
		switch key {

		case "message":
			if v, ok := value.(string); ok {
				res.Message = &v
			}

		case "err_code":
			if v, ok := value.(string); ok {
				res.ErrCode = &v
			}

		case "err_msg":
			res.ErrMsg = value

		case "data":
			res.Data = value

		case "errors":
			res.Errors = value

		case "pagination":
			if v, ok := value.(map[string]any); ok {
				res.Pagination = v
			}
		}
	}

	rw.Header().Set("Content-Type", "application/json")

	if options.StatCode >= 400 && options.StatCode <= 500 {
		rw.WriteHeader(int(options.StatCode))
		parser.Encode(rw, res)
	} else {
		parser.Encode(rw, res)
	}
}
