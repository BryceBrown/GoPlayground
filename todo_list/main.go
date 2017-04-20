package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("Starting Server...")

	r := mux.NewRouter()
	r.HandleFunc("/", index)
	http.ListenAndServe(":8080", r)
	// s := &http.Server{
	// 	Addr: ":8080",
	// 	Handler: hndl,
	// }
	// s.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		fmt.Errorf("Oh Noez!! %v", err)
	}
}