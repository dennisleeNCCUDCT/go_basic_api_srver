package controller

import (
	"encoding/json"
	"net/http"
)

func create() http.HandlerFunc{
return func(w http.ResponseWriter, r *http.Request){
	if r.Method==http.MethodPost{
//1.take some data
data :=view.PostRequest{}
json.NewEncoder(r.Body).Decode(&data)
//2.save it in mySQL

//
if err :=model.CreateTodo(data.name,data.todo);err!=nil{
	w.Write([]byte("Some error"))
	return
}
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(data)

	}
}

}