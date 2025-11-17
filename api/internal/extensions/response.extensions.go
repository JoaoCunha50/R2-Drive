package extensions

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
	Title string `json:"title"`
}

func SuccessResponse(data interface{}) Response {
	return Response{
		Success: true,
		Message: "SUCCESS",
		Data: data,
	}
}

func ErrorResponse(message string) Response {
	return Response{
		Success: false,
		Message: message,
	}
}