package textrn

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "encoding/json"
    "time"
    "os"
)

var (
    // Environment variables, will need to migrate logic to outside of package
    connect_sid = os.Getenv("TEXTRN_CONNECT_SID")
    username = os.Getenv("TEXTRN_USERNAME")
    phone = os.Getenv("TEXTRN_PHONE")

    // Base URL for TextNow service
    base_url = "https://www.textnow.com/api/users/"

    // Client used for caching TCP connections
    client = &http.Client{}
)

type Client struct {
    ConnectSid  string
    Username    string
}


func GetMessages() {
    // Create client and request objects
    req, err := http.NewRequest("GET", base_url + username + "/messages", nil)
    if err != nil {
        fmt.Println("error")
        return
    }
    // Add cookie with identifier
    req.Header.Set("Cookie", "connect.sid=" + connect_sid)

    res, err := client.Do(req)
    defer res.Body.Close()
    data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
    // Print response
    fmt.Println(string(data))
}

type SendMessageBody struct {
    PhoneNumber         string  `json:"contact_value"`
    ContactType         int     `json:"contact_type"`
    Message             string  `json:"message"`
    Read                int     `json:"read"`
    MessageDirection    int     `json:"message_direction"`
    MessageType         int     `json:"message_type"`
    FromName            string  `json:"from_name"`
    HasVideo            bool    `json:"has_video"`
    New                 bool    `json:"new"`
    Date                string  `json:"date"`
}

type MessageId struct {
    Id  string  `json:"id"`
}

/*
    Sends an SMS message to the specified number
*/
func (c Client) SendMessage(number, message string) string {
    // Construct message body
    reqBody := SendMessageBody{
        number,
        1,
        message,
        1,
        2,
        1,
        c.Username,
        false,
        true,
        time.Now().Format(time.RFC3339),
    }

    // Convert body to JSON
    bytes, err := json.Marshal(reqBody)
    if err != nil {
        fmt.Println("err")
        return "err"
    }

    // Create HTTP POST request object
    req, err := http.NewRequest("POST", base_url + username + "/messages", strings.NewReader(string(bytes)))
    if err != nil {
        fmt.Println("err")
        return "err"
    }

    // Set appropriate headers
    // The Cookie header containing the connect.sid is required for authentication with the service
    req.Header.Set("Cookie", "connect.sid=" + c.ConnectSid)
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Host", "www.textnow.com")
    req.Header.Set("Connection", "keep-alive")

    // Perform HTTP request
    res, err := client.Do(req)
    if err != nil {
        fmt.Println("err")
        return "err"
    }

    // Read response data
    defer res.Body.Close()
    data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println("err")
        return "err"
    }

    return string(data)
}

