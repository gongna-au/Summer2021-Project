package di

import (
	"net/http"

	"github.com/Summer2021-Project/di"
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
