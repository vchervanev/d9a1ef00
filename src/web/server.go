package web

import (
	"log"

	"github.com/valyala/fasthttp"
)

// Server to handle incoming requests
type Server struct {
	Address string
}

// Start method starts http listener
func (s Server) Start() {
	log.Printf("Starting at %s\n", s.Address)
	if err := fasthttp.ListenAndServe(s.Address, s.requestHandler); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}

// TODO move me
func eql(bytes1, bytes2 []byte) bool {
	if len(bytes1) != len(bytes2) {
		return false
	}
	for i, b := range bytes1 {
		if b != bytes2[i] {
			return false
		}
	}
	return true
}

var usersNew = []byte("/users/new")

func (s Server) requestHandler(ctx *fasthttp.RequestCtx) {
	log.Printf("%s %s", ctx.Method(), ctx.Path())
	if eql(ctx.Path(), usersNew) {
		log.Printf("New user: %s", string(ctx.PostBody()))
	} else {
		log.Println("unknowns path")
	}
}
