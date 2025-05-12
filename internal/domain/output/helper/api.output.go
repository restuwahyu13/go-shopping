package hopt

type (
	Error struct {
		Name    string `json:"name,omitempty"`
		Message string `json:"message"`
		Code    int    `json:"code,omitempty"`
		Stack   any    `json:"stack,omitempty"`
	}

	Response struct {
		StatCode   float64        `json:"stat_code,omitempty"`
		Message    any            `json:"message,omitempty"`
		ErrCode    any            `json:"err_code,omitempty"`
		ErrMsg     any            `json:"err_msg,omitempty"`
		Data       any            `json:"data,omitempty"`
		Errors     any            `json:"errors,omitempty"`
		Pagination map[string]any `json:"pagination,omitempty"`
	}
)
