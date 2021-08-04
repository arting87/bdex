package controller

type GetOneNewsData struct {
	Id int64 `json:"id"`
}

type GetTypeNewsData struct {
	TypeCode int `json:"typeCode"`
}

type InsertNewsData struct {
	Title    string `json:"title"`
	TypeCode int    `json:"typeCode"`
	Content  string `json:"content"`
}

type UpdataNewsData struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	TypeCode int    `json:"typeCode"`
	Content  string `json:"content"`
}

type DeleteNewsData struct {
	Id int64 `json:"id"`
}

type NewsError struct {
	Error  string `json:"error"`
	Status bool   `json:"status"`
}
type NewsSuccess struct {
	Status bool `json:"status"`
}
