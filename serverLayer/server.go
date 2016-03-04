package main

import (
  "net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("../clientLayer")))
	http.ListenAndServe(":8000", nil)
  log.Println("Server listening on port 8000")
}