package main

import (
	"flag"

	"./web"
)

func main() {
	address := flag.String("a", "0.0.0.0:80", "web server binding Address")
	path := flag.String("z", "/tmp/data/data.zip", "path to Zipped data")

	flag.Parse()

	server := web.Server{Address: *address}

	server.Preload(*path)
	server.Start()
}
