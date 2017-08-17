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

func (s Server) requestHandler(ctx *fasthttp.RequestCtx) {
	log.Printf("%s %s", ctx.Method(), ctx.Path())
	// fmt.Fprintf(ctx, "Hello, world! Requested path is %q", ctx.Path())
}
