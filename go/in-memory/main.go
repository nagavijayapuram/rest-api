package main

import (
  "github.com/gorilla/mux"
  "encoding/json"
  "net/http"
  "strconv"
  "log"
  "fmt"
)

// -----
// Post
// -----
type Post struct {
  Id string `json:"id"`
  Title string `json:"title"`
  Body string `json:"body"`
}

// --------------------
// collection of posts
// --------------------
var posts = make([]Post, 0)

// ----------
// get posts
// ----------
func getPosts(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(posts)
}

// ---
// id
// ---
var _id int
func id() string {
  _id = _id + 1
  return strconv.Itoa(_id)
}

// ------------
// create post
// ------------
func createPost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var post Post
  _ = json.NewDecoder(r.Body).Decode(&post)
  post.Id = id()
  posts = append(posts, post)
  json.NewEncoder(w).Encode(&post)
}

// ---------
// get post
// ---------
func getPost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for _, item := range posts {
    if item.Id == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Post{})
}

// ------------
// update post
// ------------
func updatePost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range posts {
    if item.Id == params["id"] {
      posts = append(posts[:index], posts[index+1:]...)
      var post Post
      _ = json.NewDecoder(r.Body).Decode(&post)
      post.Id = params["id"]
      posts = append(posts, post)
      json.NewEncoder(w).Encode(&post)
      return
    }
  }
  json.NewEncoder(w).Encode(posts)
}

// ------------
// delete post
// ------------
func deletePost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range posts {
    if item.Id == params["id"] {
      posts = append(posts[:index], posts[index+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(posts)
}

// -----
// main
// -----
func main() {
  router := mux.NewRouter()

  // seed a post
  posts = append(posts, Post {
    Id: id(),
    Title: "My first post",
    Body: "This is the content of my first post",
  })

  router.HandleFunc("/posts", getPosts).Methods("GET")
  router.HandleFunc("/posts", createPost).Methods("POST")
  router.HandleFunc("/posts/{id}", getPost).Methods("GET")
  router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
  router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

  port := "8000"
  fmt.Println()
  log.Print("Listening on port ", port)
  http.ListenAndServe(":" + port, router)
}
