package main

import (
	"net/http"
)


type Response struct {//定義reponse的status code
	Code int    `json:"code"`//code 為int
	Body string `json:"body"`//body為string
}

func main(){
mux:=controller.Register()//create a new server variable

http.ListenAndServe("localhost:3000",mux)//開始營運mux server(在localhost 3000port上)
}