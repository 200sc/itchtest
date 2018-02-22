package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/200sc/itchtest/iserve"
)

var (
	port = flag.String("port", ":80", "")
)

// Todo: differ responses based on POST or GET
// Todo: handle arguments following question mark

func main() {
	err := http.ListenAndServe(*port, iserve.New())
	if err != nil {
		fmt.Println(err)
	}
}
