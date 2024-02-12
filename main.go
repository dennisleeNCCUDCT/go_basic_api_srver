package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type Response struct {
	Code int    `json:"code"`
	Body string `json:"body"`
}

func main(){
mux:=http.NewServeMux()//create a new server variable
mux.HandleFunc("/dennis",func(w http.ResponseWriter, r *http.Request){//mux server中放置一個function(handle function),"/dennis"為local host中的url路徑,後面加入一個新function:w寫入response,r為pointer,指向request的記憶體位址
	fmt.Println("request rcevied")
	w.Write([]byte("hello world"))//用byte的方式寫出response(string)
	if r.Method == http.MethodGet {
		data := Response{
			Code: http.StatusOK,
			Body: "Dennis rules",
		}
		// Create a new JSON encoder and encode the data to the response writer
		json.NewEncoder(w).Encode(data)
	}
})
http.ListenAndServe("localhost:3000",mux)//開始營運mux server(在localhost 3000port上)
}