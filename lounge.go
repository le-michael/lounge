package lounge

import (
    "fmt"
    "reflect"
    "net/http"
    "log"

    "github.com/gorilla/websocket"
)

/// Constants ///

var (
    upgrader = websocket.Upgrader {
        ReadBufferSize: 1024,
        WriteBufferSize: 1024,
        CheckOrigin: func(r *http.Request) bool { return true },
    }
)

/// Functions ///

func isFunc(fn interface{}) bool {
	return reflect.TypeOf(fn).Kind() == reflect.Func
}

//Lounge struct will store all information about the lounge.
type Lounge struct {
	boundedFuncs map[string]interface{}
}

//NewLounge will create and return a pointer to a new lounge
func NewLounge() *Lounge {
	bf := make(map[string]interface{})
	return &Lounge{bf}
}

//ListenFor is used to bind a function to a key word.
func (l *Lounge) ListenFor(key string, fn interface{}) {
    bf := l.boundedFuncs

    if _,ok := bf[key]; ok {
		panic(fmt.Sprintf("already listening for key: %v", key))
	}

	if !isFunc(fn) {
		panic(fmt.Sprintf("gave type %v looking for type func", reflect.TypeOf(fn)))
	}

	bf[key] = fn
}

//Execute will run a specified function given a key word.
func (l *Lounge) Execute(key string, args ...interface{}) {
	fn, ok := l.boundedFuncs[key]
	if !ok {
		panic(fmt.Sprintf("no function bounded to key: %v", key))
	}
	fn.(func(...interface{}))(args...)
}

//HandleConnection will generate a new client on connection.
func (l *Lounge) HandleConnection(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    client := NewClient(conn)
    log.Println("Client Connected", client)
    go client.Listen()
}
