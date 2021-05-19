package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "encoding/json"
    "time"
    "os"
    "flag"
)

var (
    connect_sid = os.Getenv("TEXTRN_CONNECT_SID")
    base_url = "https://www.textnow.com/api/users/"
    username = os.Getenv("TEXTRN_USERNAME")
    phone = os.Getenv("TEXTRN_PHONE")
    client = &http.Client{} // Client used for caching TCP connections
)

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

func SendMessage(number, message string) {
    reqBody := SendMessageBody{
        number,
        1,
        message,
        1,
        2,
        1,
        username,
        false,
        true,
        time.Now().Format(time.RFC3339),
    }
    bytes, err := json.Marshal(reqBody)
    if err != nil {
        fmt.Println("err")
        return
    }
    req, err := http.NewRequest("POST", base_url + username + "/messages", strings.NewReader(string(bytes)))
    if err != nil {
        fmt.Println("err")
        return
    }
    req.Header.Set("Cookie", "connect.sid=" + connect_sid)
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Host", "www.textnow.com")
    //req.Header.Set("Accept", "*/*")
    //req.Header.Set("Accept-Encoding", "gzip, deflate, br")
    req.Header.Set("Connection", "keep-alive")

    res, err := client.Do(req)
    if err != nil {
        fmt.Println("err")
        return
    }
    defer res.Body.Close()
    data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println("err")
        return
    }
    // Print response (id)
    fmt.Println(string(data))
}


func main() {

    msg := flag.String("msg", "", "A string.")
    flag.Parse()
    if *msg == "" {
        fmt.Println("No message passed in")
        return
    }
    SendMessage(phone, *msg)
}

