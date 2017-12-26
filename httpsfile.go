package main

import (
	"fmt"
	"net/http"
)

func main() {
	h := http.FileServer(http.Dir("."))
	err := http.ListenAndServeTLS(":8001", "rui.crt", "rui.key", h)
	//http.ListenAndServe(fmt.Sprintf(":%d", SERVER_PORT), nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("exit")
}
