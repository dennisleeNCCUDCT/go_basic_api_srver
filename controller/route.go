package controller

import (

	//"fmt"
	"net/http"
)
type Response struct {
	Code int    `json:"code"`
	Body string `json:"body"`
}
// Register registers handlers to a new ServeMux and returns it
func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/dennis", reply())
	return mux 
}