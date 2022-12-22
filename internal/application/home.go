package application

type GetHomeResponse struct {
	Message string `json:"message"`
}

func GetHome() GetHomeResponse {
	return GetHomeResponse{Message: "Welcome Home"}
}
