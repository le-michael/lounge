package lounge

import (
    "fmt"
    "github.com/gorilla/websocket"
)

type jsonmap map[string]interface{}

type Client struct {
    conn *websocket.Conn
    lounge *Lounge
}

func NewClient(conn *websocket.Conn, lounge *Lounge) *Client{
    return &Client{conn, lounge}
}

func (c *Client) Listen() {
    for {
        _, req, err := c.conn.ReadMessage()
        if err != nil {
            fmt.Println(err)
            break
        }

        mssg := JsonToMap(c, req)
        fmt.Println(mssg)
        c.lounge.Execute(mssg["task"].(string), mssg)
    }
}

func (c *Client) Send(res string) {
    c.conn.WriteMessage(websocket.TextMessage, []byte(res))
}
