package routes

import (
    "net/http"
    "github.com/gorilla/mux"
    "golang-blog-api/controllers"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/login", controllers.Login).Methods("POST")
    router.HandleFunc("/posts", controllers.GetPosts).Methods("GET")

    return router
}
