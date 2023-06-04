package helper

func ResponseFormat(code int, msg string, data any) (int, map[string]any) {
	result := map[string]any{}
	result["code"] = code
	result["message"] = msg

	if data != nil {
		result["data"] = data
	}

	return code, result
}
