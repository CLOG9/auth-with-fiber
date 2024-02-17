package shared

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type GlobalHandlerResp struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func GlobErrResp(message string) GlobalErrorHandlerResp {
	res := GlobalErrorHandlerResp{
		Success: false,
		Message: message,
	}
	return res
}
func GlobResp(data interface{}) GlobalHandlerResp {
	res := GlobalHandlerResp{
		Success: true,
		Data:    data,
	}
	return res
}
