package models

type JSONUserResponse struct {
	Data    interface{} `json:"data"`
	Sign    string      `json:"sign"`
	Message string      `json:"message"`
	IsOk    bool        `json:"isOk"`
}
type JSONGetUserResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	IsOk    bool        `json:"isOk"`
}

type JSONBookResponse struct {
	Data    Data   `json:"data"`
	Message string `json:"message"`
	IsOk    bool   `json:"isOk"`
}
type JSONGetBooksResponse struct {
	Data    []GetBookModel `json:"data"`
	Message string         `json:"message"`
	IsOk    bool           `json:"isOk"`
}
type JSONUpdateBookResponse struct {
	Data    GetBookModel `json:"data"`
	Message string       `json:"message"`
	IsOk    bool         `json:"isOk"`
}
type JSONDeleteBookResponse struct {
	Data    string `json:"data"`
	Message string `json:"message"`
	IsOk    bool   `json:"isOk"`
}
type Data struct {
	Status uint8
	Book   Book
}
type JSONErrorResponse struct {
	Error string `json:"error"`
}
