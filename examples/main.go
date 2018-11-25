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
    l.ListenFor("hello", func(data lounge.JsonMap){

        client := data["client"].(*lounge.Client)

        fmt.Println(data)
        client.Send("Hello There")
    })

    http.HandleFunc("/ws", func(w  http.ResponseWriter, r *http.Request) {
        l.HandleConnection(w, r)
    })

    http.ListenAndServe(port, nil)
}
