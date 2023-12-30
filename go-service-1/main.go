package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	hostname := os.Getenv("HOSTNAME")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!"+" -container:"+hostname)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("listen error: ", err)
	}
}
