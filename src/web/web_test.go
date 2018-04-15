package web

import (
	"net/http"
	_ "net/http/pprof"
	"testing"
)

func TestPreload(t *testing.T) {
	server := Server{}
	server.Preload("../../data/data_full.zip")
	http.ListenAndServe(":9090", http.DefaultServeMux)
}
