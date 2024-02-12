package controller

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Body string `json:"body"`
}


func reply() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))

		if r.Method == http.MethodGet {
			data := Response{
				Code: http.StatusOK,
				Body: "Dennis rules",
			}
			// Create a new JSON encoder and encode the data to the response writer
			json.NewEncoder(w).Encode(data)
		}
	}
}