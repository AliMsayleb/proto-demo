package server

import (
	"fmt"
	"net/http"
	"sef-demo/services"
)

func StartHTTP() {
	http.HandleFunc("/user", services.UserHTTPService{}.GetUser)
	fmt.Println("Listening on port 7070")
	go func() {
		http.ListenAndServe(":7070", nil)
	}()
}
