package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-martini/martini"
	"github.com/gorilla/feeds"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	//"github.com/martini-contrib/gorelic"
)

var authorName = "George Moura"
var authorEmail = "gwmoura@gmail.com"
var domain = "georgemoura.com.br"
var hasString = template.FuncMap{"hasString": func(word string, prefix string) bool {
	return strings.Contains(word, prefix)
}}

// AssetsHost - var to set host fo assets site
var AssetsHost = template.FuncMap{"assetsHost": func() string { return "https://storage.googleapis.com/george-moura-site.appspot.com" }}

// Host - var to set site url
var Host = template.FuncMap{"host": func() string { return domain }}

// Env - app env
var Env = template.FuncMap{"env": func() string { return env() }}

var inProduction = template.FuncMap{"inProduction": func() bool { return env() == "production" }}

// GaAccount - return Google Analytics ID
var GaAccount = template.FuncMap{"gaAccount": func() string {
	if env() == "production" {
		return "UA-38960674-1"
	}

	return ""
}}

func env() string {
	return os.Getenv("ENV")
}

// MyInstance - Martini innstace used on the app
var MyInstance *martini.ClassicMartini

var feed = &feeds.Feed{
	Title:       "George Moura Blog",
	Link:        &feeds.Link{Href: "//" + domain},
	Description: "Programador Web e Mobile ",
	Author:      &feeds.Author{authorName, authorEmail},
	Created:     time.Now(),
}

func main() {
	start()
}

func setUp() {
	MyInstance = martini.Classic()
	MyInstance.Use(gzip.All())
	MyInstance.Use(render.Renderer(render.Options{
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Funcs:      []template.FuncMap{AssetsHost, Host, hasString, Env, GaAccount, inProduction},
	}))
	MyInstance.Use(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Cache-Control", "public, max-age=3600")
	})

	/*
		gorelic.InitNewrelicAgent("ef64d80f06826b61c849b959f48b9c2a52dc4ac8", "George Moura Site", true)
		MyInstance.Use(gorelic.Handler)
	*/
	for _, post := range getPosts() {
		feed.Add(
			&feeds.Item{
				Title:       post.Title,
				Link:        &feeds.Link{Href: "http://" + domain + "" + post.Link},
				Description: post.Description,
				Author:      &feeds.Author{authorName, authorEmail},
				Created:     post.CreatedAt,
			})
	}
}

func setRoutes() {
	MyInstance.Get("/", index)
	MyInstance.Get("/blog", blog)
	MyInstance.Get("/euli", euli)
	MyInstance.Get("/euli/:bookname", getBook)
	MyInstance.Get("/contatos", contatos)
	MyInstance.Get("/curriculo/:locale", showCVPage)
	MyInstance.Get("/curriculo", showCVPage)
	MyInstance.Get("/cv/:locale", showCVPage)
	MyInstance.Get("/cv", showCVPage)
	MyInstance.Get("/feed", pageFeed)
	MyInstance.Get("/natal", natal)
	MyInstance.Get("/:postname", getPost)
	MyInstance.Get("/.well-known/acme-challenge/:code", letsencrypt)
	MyInstance.NotFound(func() string {
		return "Page not found!"
	})
}

func start() {
	setUp()
	setRoutes()

	http.Handle("/", MyInstance)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func age(birthday time.Time) int {
	now := time.Now()
	years := now.Year() - birthday.Year()
	if now.YearDay() < birthday.YearDay() {
		years--
	}
	return years
}

func index(params martini.Params, r render.Render) {
	s := struct {
		Title       string
		Description string
		Keywords    string
		Posts       []Post
	}{posts[0].Title, posts[0].Description, posts[0].Keywords, getPosts()}
	r.HTML(200, "home", s)
}

func blog(params martini.Params, r render.Render) {
	s := struct {
		Title       string
		Description string
		Keywords    string
		Posts       []Post
	}{posts[0].Title, posts[0].Description, posts[0].Keywords, getPosts()}

	r.HTML(200, "blog", s)
}

func contatos(params martini.Params, r render.Render) {
	s := struct {
		Title       string
		Description string
		Keywords    string
	}{posts[0].Title, posts[0].Description, posts[0].Keywords}
	r.HTML(200, "contact", s)
}

func letsencrypt(params martini.Params, r render.Render) (int, string) {
	data, err := ioutil.ReadFile("certificates/" + params["code"])
	if err != nil {
		return 404, "Post not found!"
	}
	return 200, string(data)
}

func pageFeed(params martini.Params, r render.Render) {
	rss := getFeed()
	r.Text(200, rss)
}

func natal(params martini.Params, r render.Render) {
	s := struct {
		Title       string
		Description string
		Keywords    string
	}{"Página para o natal da família", "Página para o natal da família", "natal, família"}
	r.HTML(200, "natal", s)
}

func getFeed() string {
	rss, err := feed.ToRss()
	if err != nil {
		fmt.Print("Errors!")
	}
	return rss
}

func getPosts() []Post {
	var p []Post
	for _, post := range posts {
		if post.PostType != "page" {
			p = append(p, post)
		}
	}
	return p
}

func getPostByName(name string) Post {
	var p Post
	for _, post := range posts {
		if post.FriendlyId == name {
			p = post
		}
	}

	return p
}

func getPost(params martini.Params, r render.Render) {
	post := getPostByName(params["postname"])
	if post.FriendlyId != "" {
		r.HTML(200, "posts/"+post.FriendlyId, post)
	} else {
		r.Text(404, "Not found")
	}
}

func showCVPage(params martini.Params, r render.Render) {
	locale := params["locale"]
	templateName := "cv/br"
	fmt.Println("Locale", locale)
	if locale != "" {
		templateName = "cv/" + locale
	}
	yearsOld := age(time.Date(1990, 5, 14, 0, 0, 0, 0, time.UTC))

	s := struct {
		YearsOld    int
		Title       string
		Description string
		Keywords    string
	}{yearsOld, posts[1].Title, posts[1].Description, posts[1].Keywords}
	r.HTML(200, templateName, s)
}

func euli(params martini.Params, r render.Render) {
	s := struct {
		Title       string
		Description string
		Keywords    string
	}{posts[0].Title, posts[0].Description, posts[0].Keywords}

	r.HTML(200, "euli", s)
}

func getBook(params martini.Params, r render.Render) {
	post := getBookByName(params["bookname"])
	if post.FriendlyId != "" {
		r.HTML(200, "euli/"+post.FriendlyId, post)
	} else {
		r.Text(404, "Not found")
	}
}

func getBookByName(name string) Post {
	var p Post
	for _, post := range books {
		if post.FriendlyId == name {
			p = post
		}
	}

	return p
}
