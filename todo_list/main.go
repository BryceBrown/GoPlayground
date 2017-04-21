package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"encoding/json"
)

func main(){
	fmt.Println("Starting Server...")

	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/TodoList/", todoList)

	http.ListenAndServe(":8080", r)
	// s := &http.Server{
	// 	Addr: ":8080",
	// 	Handler: hndl,
	// }
	// s.ListenAndServe()
}

func todoList(w http.ResponseWriter, r *http.Request){

}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		fmt.Errorf("Oh Noez!! %v", err)
	}
}