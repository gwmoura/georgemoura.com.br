package site

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

func init() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Get("/hello", func(r render.Render) {
		r.HTML(200, "hello", "jeremy")
	})
	http.Handle("/", m)
}
