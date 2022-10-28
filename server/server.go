package server

import "net/http"

type Country struct {
	Name     string
	Language string
}

var countries []Country = []Country{
	{"Argentina", "Spanish"},
	{"Brasil", "Portuguese"},
	{"Chile", "Spanish"},
	{"Colombia", "Spanish"},
	{"Ecuador", "Spanish"},
	{"Paraguay", "Spanish"},
	{"Peru", "Spanish"},
	{"Uruguay", "Spanish"},
}

func New(addr string) *http.Server {
	initRoutes()
	return &http.Server{
		Addr: addr,
	}
}
