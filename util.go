package lounge

import (
    "encoding/json"
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
