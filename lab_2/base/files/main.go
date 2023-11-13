package main

import (
  "log"

  "net/http"
)

func HelloWorldServer(writer http.ResponseWriter, request *http.Request) {
  writer.Header().Set("Content-Type", "text/plain")
  writer.Write([]byte("My secure web server"))
}

func main() {
  // get our ca and server certificate
  http.HandleFunc("/", HelloWorldServer)
  err := http.ListenAndServeTLS(":443", "domain.crt", "domain.key", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
