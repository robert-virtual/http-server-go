package main

import (
	"fmt"
	"net/http"
)

func posts(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		getPosts(res, req)
	case "POST":
		createPost(res, req)
  default:
    notimplemented(res,req)
	}
}

func notimplemented(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "%v posts method not implemented",req.Method)
}

func getPosts(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "get posts")
}

func createPost(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "create posts")
}
