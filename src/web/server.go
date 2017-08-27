package web

import (
	"log"

	"github.com/valyala/fasthttp"

	"../tools/bytes"

	"../entity/model"
	"../entity/model/location"
	"../entity/model/user"
	"../entity/model/visit"
	"../entity/storage"
	"../zip"
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

func parse(record []byte, filter byte) {
	switch filter {
	case 'u':
		db.Add(user.BuildUser(record))
	case 'l':
		db.Add(location.BuildLocation(record))
	case 'v':
		db.Add(visit.BuildVisit(record))
	default:
		log.Fatalf("Unknowns type %v", filter)
	}
}

// load zipped data
func (s Server) Preload(path string) {
	log.Printf("Preloading from %s\n", path)
	zip.LoadObjects(path, []byte("ulv"), parse)
	log.Println(db.Info())
}

var db = storage.MemoryServiceFactory.CreateMemoryService([]string{"user", "location", "visit"}, []int{1000100, 761400, 10000800})

var usersNew = []byte("/users/new")
var locationsNew = []byte("/locations/new")
var visitsNew = []byte("/visits/new")

var usersGet = []byte("/users/")
var locationsGet = []byte("/locations/")
var visitsGet = []byte("/visits/")

var emptyResponse = []byte("{}")

var postMethod = []byte("POST")

var nullValue = []byte("\": null")

func (s Server) requestHandler(ctx *fasthttp.RequestCtx) {
	// TODO only in dev
	// log.Printf("%s %s", ctx.Method(), ctx.Path())

	if bytes.Eql(postMethod, ctx.Method()) {
		if bytes.IndexOfSubarray(ctx.PostBody(), nullValue) != -1 {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			ctx.Write(emptyResponse)
		} else if bytes.Eql(ctx.Path(), usersNew) {
			userRecord := user.BuildUser(ctx.PostBody())
			db.Add(userRecord)
			ctx.Write(emptyResponse)
		} else if bytes.Eql(ctx.Path(), locationsNew) {
			locationRecord := location.BuildLocation(ctx.PostBody())
			db.Add(locationRecord)
			ctx.Write(emptyResponse)
		} else if bytes.Eql(ctx.Path(), visitsNew) {
			visitRecord := visit.BuildVisit(ctx.PostBody())
			db.Add(visitRecord)
			ctx.Write(emptyResponse)
		} else if bytes.StartsWith(ctx.Path(), usersGet) || bytes.StartsWith(ctx.Path(), locationsGet) || bytes.StartsWith(ctx.Path(), visitsGet) {
			id := bytes.GetId(ctx.Path(), '/')

			record := db.Get(entityType(ctx.Path()), id)
			if record != nil {
				model.UpdateRecord(record, ctx.PostBody())
				db.Update(*record)
				ctx.Write(emptyResponse)
			} else {
				ctx.SetStatusCode(fasthttp.StatusNotFound)
			}

		} else {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
		}
	} else {
		if bytes.StartsWith(ctx.Path(), usersGet) || bytes.StartsWith(ctx.Path(), locationsGet) || bytes.StartsWith(ctx.Path(), visitsGet) {
			id := bytes.GetId(ctx.Path(), '/')
			user := db.Get(entityType(ctx.Path()), id)
			if user != nil {
				user.WriteJSON(ctx)
			} else {
				ctx.SetStatusCode(fasthttp.StatusNotFound)
			}
		} else {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
		}
	}
}

func entityType(path []byte) string {
	c := path[1]
	entityType := ""

	if c == 'u' {
		entityType = "user"
	} else if c == 'l' {
		entityType = "location"
	} else {
		entityType = "visit"
	}
	return entityType
}
