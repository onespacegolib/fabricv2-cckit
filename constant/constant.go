package constant

type (
	Json      map[string]interface{}
	ApiResult struct {
		Result      interface{} `json:"result"`
		Message     interface{} `json:"message"`
		MessageCode int         `json:"message_code"`
	}
)
