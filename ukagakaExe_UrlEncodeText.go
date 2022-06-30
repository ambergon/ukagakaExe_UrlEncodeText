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
    var searchText string = ""

    if len(args) != 0 {
        for i := 0; i < len( args ) ; i++ {
            searchText = searchText + args[i] + " "
        }
    }else{
        //searchText = "こんにちは"
        os.Exit(0)
    }

    text := url.QueryEscape( searchText )
    fmt.Print( text )
}
