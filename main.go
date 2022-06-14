package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hola Mundo")
	})
  http.HandleFunc("/posts",posts)
  port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	host := fmt.Sprintf(":%s", port)
	result := fmt.Sprintf("server runnin on port %s...", port)
	fmt.Printf(result)
	http.ListenAndServe(host, nil)
}
