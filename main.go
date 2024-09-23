package main

import (
    "log"
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("./"))
    http.Handle("/", fs)

    port := ":3000"
    log.Printf("Server running on http://localhost%s\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatal(err)
    }
}
