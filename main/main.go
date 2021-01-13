package main

import (
    "os"
    "fmt"
    "log"
    "strconv"
    "net/http"
    "github.com/liberdade-organizacao/message-queue/database"
)

func main() {
    port := os.Getenv("PORT")
    http.HandleFunc("/", index)
    http.HandleFunc("/setup", setup)
    http.HandleFunc("/create", newMessage)
    http.HandleFunc("/read", getMessages)
    log.Print("Serving at ", port)
    log.Fatal(http.ListenAndServe(":" + port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "ping: %s", database.Ping(os.Getenv("DATABASE_URL")))
}

func setup(w http.ResponseWriter, r *http.Request) {
    databaseUrl := os.Getenv("DATABASE_URL")
    oops := database.Setup(databaseUrl)
    if oops == nil {
        fmt.Fprintf(w, "ok")
    } else {
        fmt.Fprintf(w, "error: %v", oops)
    }
}

func newMessage(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    kind := r.Form["kind"][0]
    contents := r.Form["content"][0]

    databaseUrl := os.Getenv("DATABASE_URL")
    oops := database.NewMessage(databaseUrl, kind, contents)

    if oops == nil {
        fmt.Fprintf(w, "ok")
    } else {
        fmt.Fprintf(w, "error: %v", oops)
    }
}

func getMessages(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    kind := r.Form["kind"][0]
    offset, oops := strconv.Atoi(r.Form["offset"][0])
    if oops != nil {
        panic(oops)
    }

    databaseUrl := os.Getenv("DATABASE_URL")
    messages, oops := database.GetMessages(databaseUrl, kind, offset)

    if oops != nil {
        fmt.Fprintf(w, "error: %v", oops)
    } else if messages == nil {
        fmt.Fprintf(w, "error: no messages")
    } else {
        for _, message := range messages {
            fmt.Fprintf(w, "---\n")
            fmt.Fprintf(w, "id: %s\n", message["id"])
            fmt.Fprintf(w, "kind: %s\n", message["kind"])
            fmt.Fprintf(w, "content: %s\n", message["content"])
        }
        fmt.Fprintf(w, "...\n")
    }
}
