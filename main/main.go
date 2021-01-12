package main

import (
    "os"
    "fmt"
    "log"
    "net/http"
    "database/sql"
    _ "github.com/lib/pq"
)

func main() {
    port := os.Getenv("PORT")
    http.HandleFunc("/", index)
    log.Print("Serving at ", port)
    log.Fatal(http.ListenAndServe(":" + port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
    // database setup
    databaseUrl := os.Getenv("DATABASE_URL")
    db, oops := sql.Open("postgres", databaseUrl)
    if oops != nil {
        panic(oops)
    }
    defer db.Close()

    // processing request
    oops = db.Ping()

	fmt.Fprintf(w, "ping: %s", oops)
}
