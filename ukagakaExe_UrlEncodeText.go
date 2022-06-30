package main

import (
    "fmt"
    "net/url"
    "os"
    "flag"
)

func main() {

    flag.Parse()
    args := flag.Args()
    var searchText string

    if len(args) == 1 {
        searchText = args[0]
    }else{
        //searchText = "こんにちは"
        os.Exit(0)
    }

    text := url.QueryEscape( searchText )
    fmt.Println( text )
}
