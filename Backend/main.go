package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/gorilla/mux"
)

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/welcome", welcomeHandler).Methods("GET")

  fmt.Println("Server is running on port 8080")
  log.Fatal(http.ListenAndServe(":8080", r))
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Welcome to the Go backend!"))
}
