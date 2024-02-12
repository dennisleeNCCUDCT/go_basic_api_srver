package main

import (
	"fmt"
	"log"
	"net/http"
	//"github.com/go-sql-driver/mysql"
)


type Response struct {//定義reponse的status code
	Code int    `json:"code"`//code 為int
	Body string `json:"body"`//body為string
}

func main(){
mux := controller.Register()//create a new server variable
db := models.Connect()
defer db.Close()
fmt.Println("serving...")
log.Fatal(http.ListenAndServe("localhost:3000",mux))//開始營運mux server(在localhost 3000port上)
}