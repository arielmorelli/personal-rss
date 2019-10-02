package main

import (
    "fmt"

    "net/http"
    "encoding/json"
)

type PostEntry struct {
    Title string `json:"title"`
    Url string `json:"url"`
}

// TODO: remove global e use arguments
var PostEntries []PostEntry

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("hello world")
}

func returnAllPosts(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Getting all posts")
    json.NewEncoder(w).Encode(PostEntries)
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/all", returnAllPosts)
    http.ListenAndServe(":8080", nil)
}

func main() {
    PostEntries = []PostEntry{
        PostEntry{Title: "Test", Url: "https//test.com"},
        PostEntry{Title: "Google", Url: "https://google.com"},
    }
    handleRequests()
}
