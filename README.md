# Textrn

This is a Go package that encapsulates SMS messaging from a Go binary, using the TextNow service.

## Features

Currently the only supported feature is sending an SMS message to a specified number from the user's TextNow number. Accordingly, they must first register an account with TextNow.

## Usage

Note that as mentioned above, a TextNow account must first be created. Upon doing so, determine your `connect_sid` value by navigating to [here]("https://www.textnow.com/messaging" TextNow), inspecting the webpage and finding the value of the cookie `connect.sid`. 

Also remember the username that was registered with TextNow, as that will be the value of `username`.

```go
import "github.com/nathantau/textrn"
...
    client := &{connect_sid, username}
    err := client.SendMessage("<Phone Number>", "<Message Body>")
    if err == "err" {
        // An error occurred
    }
```

## Credits

Credits to [leogomezz4t's work](https://github.com/leogomezz4t/PyTextNow_API) for identifying the API endpoints, payloads, and means of authentication.

## Disclaimer

I am in no way responsible for anything you do with this code.
 
