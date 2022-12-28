package model

import "encoding/json"

type Response struct {
	User_name string `json:"user_name"`
	Task      string `json:"task"`
	Results   Result `json:"results"`
}

type Result struct {
	Payload json.RawMessage `json:"payload"`
	Results interface{}     `json:"results"`
}

func CreateResponse(task string, payload json.RawMessage, result interface{}) *Response {
	return &Response{
		User_name: "Nrvnqsr",
		Task:      task,
		Results: Result{
			Payload: payload,
			Results: result,
		},
	}

}

type ResponseSolution struct {
	Percent int `json:"percent"`
	Fails   []struct {
		OriginalResult string `json:"OriginalResult"`
		ExternalResult string `json:"ExternalResult"`
	} `json:"fails"`
}
