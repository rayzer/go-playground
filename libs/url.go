package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	urlstring := "http://10.75.144.42:30001/setup"
	u, err := url.Parse(urlstring)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Hostname())
}
