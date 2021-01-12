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
    http.HandleFunc("/setup", setup)
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

func setup(w http.ResponseWriter, r *http.Request) {
    databaseUrl := os.Getenv("DATABASE_URL")
    db, oops := sql.Open("postgres", databaseUrl)
    if oops != nil {
        panic(oops)
    }
    defer db.Close()

    _, oops = db.Exec(`
        CREATE TABLE IF NOT EXISTS messages (
            id SERIAL PRIMARY KEY,
            kind VARCHAR(32) NOT NULL,
            content VARCHAR(2048) NOT NULL,
            inclusion TIMESTAMP WITH TIME ZONE NOT NULL
        );
    `)
    if oops == nil {
        fmt.Fprintf(w, "ok")
    } else {
        fmt.Fprintf(w, "error: %v", oops)
    }
}
