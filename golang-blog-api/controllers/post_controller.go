package controllers

import "net/http"

func GetPosts(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Get Posts Endpoint"))
}
