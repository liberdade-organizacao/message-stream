package database

import (
    "strconv"
    "database/sql"
    _ "github.com/lib/pq"
)

func Ping(databaseUrl string) error {
    db, oops := sql.Open("postgres", databaseUrl)
    if oops != nil {
        panic(oops)
    }
    defer db.Close()
    return db.Ping()
}

func Setup(databaseUrl string) error {
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
        )
    ;`)

    return oops
}

func NewMessage(databaseUrl, kind, contents string) error {
    db, oops := sql.Open("postgres", databaseUrl)
    if oops != nil {
        panic(oops)
    }
    defer db.Close()

    _, oops = db.Exec(`
        INSERT INTO messages(kind,content,inclusion)
        VALUES ($1,$2,CURRENT_TIMESTAMP)
        RETURNING *
    ;`, kind, contents)

    return oops
}

func GetMessages(databaseUrl, kind string, offset int) ([]map[string]string, error) {
    db, oops := sql.Open("postgres", databaseUrl)
    if oops != nil {
        return nil, oops
    }
    defer db.Close()

    rows, oops := db.Query(`
        SELECT * FROM messages
        WHERE kind=$1
    ;`, kind)
    if oops != nil {
        return nil, oops
    }
    defer rows.Close()
    outlet := make([]map[string]string, 0)
    for rows.Next() {
        var id int
        var rowKind string
        var rowContent string
        var rowDate string
        rows.Scan(&id, &rowKind, &rowContent, &rowDate)

        message := make(map[string]string)
        message["id"] = strconv.Itoa(id)
        message["kind"] = rowKind
        message["content"] = rowContent

        outlet = append(outlet, message)
    }

    return outlet, nil
}
