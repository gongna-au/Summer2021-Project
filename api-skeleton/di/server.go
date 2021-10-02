package di

import (
	"github.com/Summer2021-Project/di"
	"net/http"
)

func init() {
	obj := di.Object{
		Name: "server",
		New: func() (i interface{}, e error) {
			return &http.Server{}, nil
		},
	}
	if err := di.Provide(&obj); err != nil {
		panic(err)
	}
}

func Server() (s *http.Server) {
	if err := di.Populate("server", &s); err != nil {
		panic(err)
	}
	return
}
