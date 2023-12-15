package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloworld)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("에러 발생")
		panic(err)
		return
	}
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, world")
}
