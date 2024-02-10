package main

import (
  "net/http"
  "log"
  "github.com/leonhsi/music-player/db/sqlc"
)

func main() {

  http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
    log.Println("Hello World")
  }) 

  http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
    log.Println("Goodbye World")
  }) 

  http.ListenAndServe(":1111", nil)
}
