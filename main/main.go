package main

import (
    "os"
    "fmt"
    "log"
    "net/http"
)

func main() {
    port := os.Getenv("PORT")
    http.HandleFunc("/", index)
    log.Print("Serving at ", port)
    log.Fatal(http.ListenAndServe(":" + port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!")
}
