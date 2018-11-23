package main

import (
    "fmt"
    "net/http"

    "github.com/le-michael/lounge"
)

const (
    port = ":3000"
)



func main() {

    l := lounge.NewLounge()
    l.ListenFor("hello", func(args ...interface{}){
        fmt.Println(args)
        fmt.Println("hello",args[0])
    })

    http.HandleFunc("/ws", func(w  http.ResponseWriter, r *http.Request) {
        l.HandleConnection(w, r)
    })

    http.ListenAndServe(port, nil)
}
