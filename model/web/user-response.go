package web

type UserResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Login struct {
	Token string `json:"token"`
}