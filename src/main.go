package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

// Handler function to serve files
func fileHandler(w http.ResponseWriter, r *http.Request) {
    filePath := "." + r.URL.Path
    if filePath == "./" {
        filePath = "./index.html"
    }

    file, err := os.Open(filePath)
    if err != nil {
        http.Error(w, "File not found.", http.StatusNotFound)
        return
    }
    defer file.Close()

    http.ServeFile(w, r, filePath)
}

func main() {
    port := "8080"
    http.HandleFunc("/", fileHandler)
    fmt.Printf("Starting GoServe on port %s...\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}