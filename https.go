package main

import (
	"fmt"
	"net/http"
)

const (
	SERVER_PORT       = 8080
	SERVER_DOMAIN     = "localhost"
	RESPONSE_TEMPLATE = "hello"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("hello")
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(len(RESPONSE_TEMPLATE)))
	w.Write([]byte(RESPONSE_TEMPLATE))
}

func main() {
	http.HandleFunc("/", rootHandler)
	err:=http.ListenAndServeTLS(fmt.Sprintf(":%d", SERVER_PORT), "rui.crt", "rui.key", nil)
	//http.ListenAndServe(fmt.Sprintf(":%d", SERVER_PORT), nil)
	if err!=nil {
		fmt.Println(err.Error())
	}
	
	fmt.Printf("exit")
}
