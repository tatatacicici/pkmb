package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "os"
    "time"

    _ "github.com/lib/pq"
    "github.com/gorilla/mux"
)

type User struct {
    ID        int       `json:"id"`
    Username  string    `json:"username"`
    FullName  string    `json:"full_name"`
    CreatedAt time.Time `json:"created_at"`
}

func main() {
    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        log.Fatal("DATABASE_URL not set")
    }

    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatalf("failed to open db: %v", err)
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        log.Fatalf("db ping failed: %v", err)
    }
    log.Println("users service: connected to db")

    r := mux.NewRouter()
    r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("ok"))
    }).Methods("GET")

    r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query("SELECT id, username, full_name, created_at FROM users")
        if err != nil {
            http.Error(w, "db error", 500)
            return
        }
        defer rows.Close()
        var out []User
        for rows.Next() {
            var u User
            if err := rows.Scan(&u.ID, &u.Username, &u.FullName, &u.CreatedAt); err != nil {
                http.Error(w, "scan error", 500)
                return
            }
            out = append(out, u)
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(out)
    }).Methods("GET")

    addr := ":8002"
    log.Printf("users service listening on %s", addr)
    log.Fatal(http.ListenAndServe(addr, r))
}
