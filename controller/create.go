package controller

import (
	"net/http"
)

func create() http.HandlerFunc{
return func(w http.ResponseWriter, r *http.Request){
	if r.Method==http.MethodGet{
//1.take some data
//2.save it in mySQL
if err :=model.CreateTodo();err!=nil{
	w.Write([]byte("Some error"))
	return
}

	}
}

}