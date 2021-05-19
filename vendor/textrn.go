package textrn

import (
    "fmt"
    "net/http"
)

var (
    connect_sid = "s%3A5myH69r2kQp_CUw66QzD0s_FqwGa-WOJ.zPcojx21wL0sq0JBS8JqBLbcDIEOA%2BegvJwOIwgd0n8"
    base_url = "https://www.textnow.com/api/users"
    username = "natenate1280"
)

func Messages() {
    res, err := http.Get(base_url + "/" + username + "/messages")
    fmt.Println(res)
}

func Main() {
    Messages()
}

