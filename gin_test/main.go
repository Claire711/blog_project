package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "<h1>I hope I can success</h1>")

}

func main() {
	http.HandleFunc("/", sayHello)
	fmt.Println("Starting the server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v\n", err)
		return
	}

}
