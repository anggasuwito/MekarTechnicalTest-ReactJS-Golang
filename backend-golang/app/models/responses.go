package models

//Response is a struct for wrap responses
type Response struct {
	Meta MetaComponents `json:"meta"`
	Data interface{}    `json:"data"`
}

//MetaComponents is a struct for meta components
type MetaComponents struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Records int    `json:"records"`
}
