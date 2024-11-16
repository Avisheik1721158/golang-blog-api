package main

import (
    "log"
    "net/http"
    "golang-blog-api/routes"
)

func main() {
    router := routes.SetupRoutes()
    log.Println("Starting server on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", router))
}
