package web

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"

	"./tools/bytes"

	"../entity/model/user"
	"../entity/storage"
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

var db = storage.MemoryServiceFactory.CreateMemoryService([]string{"user"})

var usersNew = []byte("/users/new")

var usersGet = []byte("/users/")

func (s Server) requestHandler(ctx *fasthttp.RequestCtx) {
	log.Printf("%s %s", ctx.Method(), ctx.Path())
	if bytes.Eql(ctx.Path(), usersNew) {
		userRecord := user.BuildUser(ctx.PostBody())
		db.Add(userRecord)
	} else if bytes.StartsWith(ctx.Path(), usersGet) {
		id := bytes.GetId(ctx.Path(), '/')
		user := db.Get("user", id)
		if user != nil {
			user.WriteJSON(ctx)
		} else {
			fmt.Fprintln(ctx, "not found")
		}
	} else {
		log.Println("unknowns path")
	}
}
