package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Running.")

	})

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPosts).Methods("POST")
	log.Println("Server listening on port: 8000")
	http.ListenAndServe(":8000", router)

}

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

//init method for initializing with one post
func init() {
	posts = []Post{Post{ID: 1, Title: "One", Text: "Hello"}}
}

//get all posts
func getPosts(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(posts)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error occured}`))
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(result)

}

//adding new posts
func addPosts(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-type", "application/json")

	var post Post
	err := json.NewDecoder(req.Body).Decode(&post)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error occured}`))
		return
	}
	res.WriteHeader(http.StatusOK)
	post.ID = len(posts) + 1
	posts = append(posts, post)
	result, err := json.Marshal(post)
	res.Write(result)

}
