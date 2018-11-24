package lounge

import (
    "fmt"
    "encoding/json"
    "github.com/gorilla/websocket"
)

type jsonmap map[string]interface{}

type Client struct {
    conn *websocket.Conn
}

func NewClient(conn *websocket.Conn) *Client{
    return &Client{conn}
}

func (c *Client) Listen() {
    for {
        _, req, err := c.conn.ReadMessage()
        if err != nil {
            fmt.Println(err)
            break
        }

        var mssg jsonmap
        json.Unmarshal(req, &mssg)
        fmt.Println(mssg)
    }
}

func (c *Client) Send(res string) {
    c.conn.WriteMessage(websocket.TextMessage, []byte(res))
}
