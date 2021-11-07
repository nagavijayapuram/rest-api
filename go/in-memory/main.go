package main

import (
  "github.com/gorilla/mux"
  "encoding/json"
  "net/http"
  "strconv"
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

// -------
// new id
// -------
var count int
func newId() string {
  count = count + 1
  return strconv.Itoa(count)
}

// ------------
// create post
// ------------
func createPost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var post Post
  _ = json.NewDecoder(r.Body).Decode(&post)
  post.Id = newId()
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
    Id: newId(),
    Title: "My first post",
    Body: "This is the content of my first post",
  })

  router.HandleFunc("/posts", getPosts).Methods("GET")
  router.HandleFunc("/posts", createPost).Methods("POST")
  router.HandleFunc("/posts/{id}", getPost).Methods("GET")
  router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
  router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

  http.ListenAndServe(":8000", router)
}
