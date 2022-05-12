package main

import (
	"fmt"
	"net/http"
	"os"

	go_web_utils "github.com/ronGeva/go_web/utils"
)

func hello_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, here are all the previously searched urls: \n")
	for currentUrl := go_web_utils.UserRequests.SeenUrls.Front(); currentUrl != nil; currentUrl = currentUrl.Next() {
		a, ok := currentUrl.Value.(*http.Request)
		if ok {
			fmt.Fprintf(w, a.RequestURI+"\n")
		} else {
			fmt.Fprintf(w, "Something weird occurred, cast to http.request failed\n")
		}
	}
}

func nested_hello_page(w http.ResponseWriter, r *http.Request) {
	go_web_utils.UserRequests.SeenUrls.PushBack(r)
	fmt.Fprintf(w, "You requested "+r.RequestURI)
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: main.exe <port>")
		return
	}
	port := args[1]
	fmt.Println("Initializing web server at port " + port)

	http.HandleFunc("/hello", hello_page)
	http.HandleFunc("/hello/", nested_hello_page)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Something went wrong :(")
	}
}
