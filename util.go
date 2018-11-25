package lounge

import (
    "fmt"
    "log"
    "encoding/json"
    "crypto/rand"
)

type JsonMap map[string]interface{}

//JsonToMap will parse a Json message from a Client and store it into a map.
func JsonToMap (client *Client, req []byte) JsonMap {
    
    // Parse the json from the client into a map
    var res JsonMap
    json.Unmarshal(req, &res)

    // Store the client making the request
    res["client"] = client

    return res
}

//GenUUID will generate and return a 128-bit uuid.
func GenUUID () string {
    b := make([]byte, 16)
    _, err := rand.Read(b)
    if err != nil {
        log.Fatal(err)
    }
    return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
