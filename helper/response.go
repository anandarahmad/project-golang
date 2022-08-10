package helper

import "strings"

type Response struct {
	Status bool        `json:"status"`
	Pesan  string      `json:"pesan"`
	Errors interface{} `json:"errors"`
	Data   interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(status bool, pesan string, data interface{}) Response {
	res := Response{
		Status: status,
		Pesan:  pesan,
		Errors: nil,
		Data:   data,
	}
	return res
}

func BuildErrResponse(pesan string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Pesan:  pesan,
		Errors: splittedError,
		Data:   data,
	}
	return res
}
