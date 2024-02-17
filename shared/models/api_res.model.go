package shared

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func GlobErrResp(message string) GlobalErrorHandlerResp {
	res := GlobalErrorHandlerResp{
		Success: false,
		Message: message,
	}
	return res
}
