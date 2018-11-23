package lounge

import (
    "fmt"
    "reflect"
)

func isFunc(fn interface{}) bool {
    return reflect.TypeOf(fn).Kind() == reflect.Func
}


type Lounge struct {
    boundedFuncs    map[string]interface{}

}

func NewLounge() *Lounge {
    bf := make(map[string]interface{})
    return &Lounge{bf}
}

func (l *Lounge) ListenFor(key string, fn interface{}) {
    if _, ok := l.boundedFuncs[key]; ok {
        panic(fmt.Sprintf("already listening for key: %v", key))
    }
    
    if !isFunc(fn) {
        panic(fmt.Sprintf("gave type %v looking for type func",reflect.TypeOf(fn)))
    }

    l.boundedFuncs[key] = fn
}

func (l *Lounge) Execute(key string, args ...interface{}){
    fn, ok := l.boundedFuncs[key] 
    if !ok {
        panic(fmt.Sprintf("no function bounded to key: %v", key))
    }
    fn.(func(...interface{}))(args...)
}
 
    

