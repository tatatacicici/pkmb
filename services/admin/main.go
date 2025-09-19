package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Admin struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
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

	// simple ping
	if err := db.Ping(); err != nil {
		log.Fatalf("db ping failed: %v", err)
	}
	log.Println("admin service: connected to db")

	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}).Methods("GET")

	r.HandleFunc("/admins", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, username, role, created_at FROM admins")
		if err != nil {
			http.Error(w, "db error", 500)
			return
		}
		defer rows.Close()
		var out []Admin
		for rows.Next() {
			var a Admin
			if err := rows.Scan(&a.ID, &a.Username, &a.Role, &a.CreatedAt); err != nil {
				http.Error(w, "scan error", 500)
				return
			}
			out = append(out, a)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(out)
	}).Methods("GET")

	addr := ":8001"
	log.Printf("admin service listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
