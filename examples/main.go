package main

import (
    "fmt"
    "github.com/le-michael/lounge"
)


func hello(t string) {
    fmt.Println("HELLO")
}

func main() {

    l := lounge.NewLounge()
    l.ListenFor("hello", func(args ...interface{}){
        fmt.Println(args)
        fmt.Println("hello",args[0])
    })
    
    l.Execute("hello","michael")
}
