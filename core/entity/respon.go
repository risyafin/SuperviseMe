package entity

type ResponsesSucces struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponSucces struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	Data    any    `json:"data"`
}

type Respons struct {
	Message string `json:"message"`
}

type ResponsesError struct {
	Error string `json:"error"`
}
