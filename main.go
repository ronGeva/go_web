package main

import (
	"fmt"
	utils "github.com/ronGeva/go_web/utils"
	"net/http"
)

func main_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, "+r.Host)
}

func main() {
	fmt.Println("hello world, " + utils.NormalizeUrl("asd"))
	http.HandleFunc("/hello", main_page)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Something went wrong :(")
	}
}
