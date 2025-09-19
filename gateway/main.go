package main

import (
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "os"
)

func newProxy(target string) *httputil.ReverseProxy {
    u, err := url.Parse(target)
    if err != nil {
        log.Fatalf("invalid proxy url: %v", err)
    }
    return httputil.NewSingleHostReverseProxy(u)
}

func main() {
    adminProxy := newProxy("http://admin:8001")
    usersProxy := newProxy("http://users:8002")

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("gateway ok"))
    })

    http.HandleFunc("/admin/", func(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = r.URL.Path[len("/admin"):] // strip prefix
        adminProxy.ServeHTTP(w, r)
    })

    http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = r.URL.Path[len("/users"):] // strip prefix
        usersProxy.ServeHTTP(w, r)
    })

    port := os.Getenv("GATEWAY_PORT")
    if port == "" {
        port = "8080"
    }
    addr := ":" + port
    log.Printf("gateway listening on %s", addr)
    log.Fatal(http.ListenAndServe(addr, nil))
}
