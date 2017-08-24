package main

import (
	"flag"

	"./web"
)

func main() {
	address := flag.String("a", "0.0.0.0:80", "web server binding address")
	flag.Parse()

	server := web.Server{Address: *address}

	server.Start()
}
