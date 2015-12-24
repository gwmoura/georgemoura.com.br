package site

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/gorilla/feeds"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	"net/http"
	"time"
	//"github.com/martini-contrib/gorelic"
)

var authorName string = "George Moura"
var authorEmail string = "gwmoura@gmail.com"
var domain string = "georgemoura.com.br"

var feed = &feeds.Feed{
	Title:       "George Moura Blog",
	Link:        &feeds.Link{Href: "http://georgemoura.com.br"},
	Description: "Programador Web e Mobile ",
	Author:      &feeds.Author{authorName, authorEmail},
	Created:     time.Now(),
}

type Post struct {
	Title       string
	Description string
	Link        string
	Keywords    string
	AuthorName  string
	AuthorEmail string
	CreatedAt   time.Time
	PostType    string
	FriendlyId  string
}

var posts = []Post{
	{
		Title:       "George Moura Blog",
		Description: "Blog de George Moura",
		Keywords:    "",
		PostType:    "page",
		FriendlyId:  "home",
	},
	{
		Title:       "George Moura | Currículo",
		Description: "Currículo George Moura",
		Keywords:    "",
		PostType:    "page",
		FriendlyId:  "cv",
	},
	{
		Description: "Estou escrevendo este post para não esquecer como configurar novas resoluções de monitores no meu ubuntu.",
		Keywords: "ubuntu, monitor",
		Title: "Adicionando resolução do monitor no ubuntu 14.04",
		Link: "/adicionando-resolucao-do-monitor-no-ubuntu-1404/",
		CreatedAt: time.Date(2015, time.December, 15, 10, 46, 0, 0, time.UTC).Local(),
		FriendlyId: "adicionando-resolucao-do-monitor-no-ubuntu-1404",
	},
	{
		Description: "Neste tutorial irei mostrar como instalar o docker no ubuntu 14.04",
		Keywords: "docker,ubuntu",
		Title: "Instalando Docker no Ubuntu 14.04",
		Link: "/instalando-docker-no-ubuntu-1404/",
		CreatedAt: time.Date(2015, time.November, 22, 17, 15,0,0, time.UTC).Local(),
		FriendlyId: "instalando-docker-no-ubuntu-1404",
	},
	{
		Description: "Neste tutorial vou mostrar como criar VMs na nuvem da Google (Google Cloud Platform). O primeiro passo é acessar: cloud.google.com e clicar em \"Sign in\" no canto superior direito. Após o login (com sua conta google) clique em \"My Console\" no canto superior direito. Ao clicar nesse link você irá para a página de projetos.",
		Keywords:    "google,cloud,vm,google cloud platform,compute engine",
		Title:       "Criando VMs na Google Cloud Platform",
		Link:        "/criando-vms-na-google-cloud-platform/",
		//Description: "<![CDATA[<p>Neste tutorial vou mostrar como criar VMs na nuvem da Google (Google Cloud Platform). O primeiro passo é acessar: cloud.google.com e clicar em \"Sign in\" no canto superior direito. Após o login (com sua conta google) clique em \"My Console\" no canto superior direito. Ao clicar nesse link você irá para a página de projetos. [&#8230;]</p> <p>The post <a rel=\"nofollow\" href=\"http://georgemoura.com.br/criando-vms-na-google-cloud-platform/\">Criando VMs na Google Cloud Platform</a> appeared first on <a rel=\"nofollow\" href=\"http://georgemoura.com.br\">George Moura</a>.</p>]]>",
		//Author:      &feeds.Author{authorName, authorEmail},
		CreatedAt:  time.Date(2015, time.March, 17, 18, 2, 37, 0, time.UTC).Local(),
		FriendlyId: "criando-vms-na-google-cloud-platform",
	},
	{
		Description: "O que é o Docker? \"An open platform for distributed applications\" - Uma plataforma aberta para aplicativos distribuídos. Essa é a definição no site, mas não fica muito clara no início, porém faz sentido depois. O Docker é uma ferramenta para distribuir aplicações em containers de forma rápida e leve, ou seja você não precisa mais",
		Keywords:    "docker,nodejs",
		Title:       "Iniciando com o Docker",
		Link:        "/iniciando-com-o-docker/",
		//Description: "<![CDATA[<p>O que é o Docker? "An open platform for distributed applications\" -&#160;Uma plataforma aberta para aplicativos distribuídos. Essa é a definição no site, mas não fica muito clara no início, porém faz sentido depois. O Docker é uma ferramenta para distribuir aplicações em containers de forma rápida e leve, ou seja você não precisa mais [&#8230;]</p> <p>The post <a rel=\"nofollow\" href=\"http://georgemoura.com.br/iniciando-com-o-docker/\">Iniciando com o Docker</a> appeared first on <a rel=\"nofollow\" href=\"http://georgemoura.com.br\">George Moura</a>.</p>]]>",
		//Author:      &feeds.Author{authorName, authorEmail},
		CreatedAt:  time.Date(2014, time.October, 13, 19, 6, 47, 0, time.UTC).Local(),
		FriendlyId: "iniciando-com-o-docker",
	},
	{
		Description: "Atenção faça os testes usando uma máquina virtual ou um pc que não tenha utilidade Aí galera venho falar para vocês sobre o Chrome OS, o sistema operacional da Google, acabei de instalar ele no meu pen-drive e já estou utilizando, a propósito estou escrevendo este post nele. Eu estava empolgado para comprar um chromebook,",
		Keywords:    "chrome os,chrome,virtualbox",
		Title:       "Chrome OS no meu pc",
		Link:        "/chrome-os-no-meu-pc/",
		//Description: "<![CDATA[<p>Atenção faça os testes usando uma máquina virtual ou um pc que não tenha utilidade Aí galera venho falar para vocês sobre o Chrome OS, o sistema operacional da Google, acabei de instalar ele no meu pen-drive e já estou utilizando, a propósito estou escrevendo este post nele. Eu estava empolgado para comprar um chromebook, [&#8230;]</p> <p>The post <a rel=\"nofollow\" href=\"http://georgemoura.com.br/chrome-os-no-meu-pc/\">Chrome OS no meu pc</a> appeared first on <a rel=\"nofollow\" href=\"http://georgemoura.com.br\">George Moura</a>.</p>]]>",
		//Author:      &feeds.Author{authorName, authorEmail},
		CreatedAt:  time.Date(2013, time.October, 12, 15, 53, 22, 0, time.UTC).Local(),
		FriendlyId: "chrome-os-no-meu-pc",
	},
	{
		Description: "Pessoal, vim deixar um post bem pequeno, mas que foi minha salvação e achei legal compartilhar. Estou em um projeto e acabei precisando fazer uma migração de dados, o problema é que meu script acabou duplicando as informações no banco de dados, imagine o problema, a sorte é que era um banco de desenvolvimento e",
		Keywords:    "mysql,banco de dados",
		Title:       "Como deletar registros duplicados no banco de dados, deixando apenas um",
		Link:        "/como-deletar-registros-duplicados-no-banco-de-dados-deixando-apenas-um/",
		//Description: "<![CDATA[<p>Pessoal, vim deixar um post bem pequeno, mas que foi minha salvação e achei legal compartilhar. Estou em um projeto e acabei precisando fazer uma migração de dados, o problema é que meu script acabou duplicando as informações no banco de dados, imagine o problema, a sorte é que era um banco de desenvolvimento e [&#8230;]</p> <p>The post <a rel=\"nofollow\" href=\"http://georgemoura.com.br/como-deletar-registros-duplicados-no-banco-de-dados-deixando-apenas-um/\">Como deletar registros duplicados no banco de dados, deixando apenas um</a> appeared first on <a rel=\"nofollow\" href=\"http://georgemoura.com.br\">George Moura</a>.</p>]]>",
		//Author:      &feeds.Author{authorName, authorEmail},
		CreatedAt:  time.Date(2013, time.October, 5, 9, 20, 19, 0, time.UTC).Local(),
		FriendlyId: "como-deletar-registros-duplicados-no-banco-de-dados-deixando-apenas-um",
	},
	{
		Description: "Meus caros estou escrevendo esse post para falar sobre uma dificuldade que tive logo quando resolvi programar para ruby &#8211; uma IDE. Comecei meus primeiros estudos utilizando o Eclipse e fazendo os teste em meu próprio notebook, mas assim que comecei a testar meu códigos em servidores me deparei com problemas: A versão do rails",
		Keywords:    "ide,ruby,rails,sublimetext,eclipse,netbeans",
		Title:       "Uma boa IDE para Ruby on Rails",
		Link:        "/uma-boa-ide-para-ruby-on-rails/",
		//Description: "<![CDATA[<p>Meus caros estou escrevendo esse post para falar sobre uma dificuldade que tive logo quando resolvi programar para ruby &#8211; uma IDE. Comecei meus primeiros estudos utilizando o Eclipse e fazendo os teste em meu próprio notebook, mas assim que comecei a testar meu códigos em servidores me deparei com problemas: A versão do rails [&#8230;]</p> <p>The post <a rel=\"nofollow\" href=\"http://georgemoura.com.br/uma-boa-ide-para-ruby-on-rails/\">Uma boa IDE para Ruby on Rails</a> appeared first on <a rel=\"nofollow\" href=\"http://georgemoura.com.br\">George Moura</a>.</p>]]>",
		//Author:      &feeds.Author{authorName, authorEmail},
		CreatedAt:  time.Date(2013, time.July, 3, 21, 23, 47, 0, time.UTC).Local(),
		FriendlyId: "uma-boa-ide-para-ruby-on-rails",
	},
}

func init(){
	start()
}

func start() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
	}))
	m.Use(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Cache-Control", "public, max-age=2592000")
	})
	m.Use(gzip.All())
	/*
	gorelic.InitNewrelicAgent("ef64d80f06826b61c849b959f48b9c2a52dc4ac8", "George Moura Site", true)
	m.Use(gorelic.Handler)
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

	m.Get("/", func(r render.Render) {
		s := struct {
			Description string
			Keywords    string
			Posts       []Post
		}{posts[0].Description, posts[0].Keywords, getPosts()}
		r.HTML(200, "home", s)
	})

	m.Get("/blog", func(r render.Render) {
		s := struct {
			Description string
			Keywords    string
			Posts       []Post
		}{posts[0].Description, posts[0].Keywords, getPosts()}

		r.HTML(200, "blog", s)
	})

	m.Get("/contatos", func(r render.Render) {
		s := struct {
			Description string
			Keywords    string
		}{posts[0].Description, posts[0].Keywords}
		r.HTML(200, "contact", s)
	})

	m.Get("/curriculo", func(r render.Render) {
		yearsOld := age(time.Date(1990, 5, 14, 0, 0, 0, 0, time.UTC))

		s := struct {
			YearsOld    int
			Description string
			Keywords    string
		}{yearsOld, posts[1].Description, posts[1].Keywords}

		r.HTML(200, "cv", s)
	})

	m.Get("/feed", func() string {
		rss := getFeed()
		return rss
	})

	m.Get("/natal", func(r render.Render) {
		s := struct {
			Description string
			Keywords 	string
		}{"Página para o natal da família", "natal, família"}
		r.HTML(200, "natal", s)
	})

	m.Get("/:postname", func(params martini.Params, r render.Render) (int, string) {
		post := getPostByName(params["postname"])
		if post.FriendlyId != "" {
			r.HTML(200, "posts/"+post.FriendlyId, post)
			return 200, ""
		}
		return 404, "Post not found!"
	})

	m.NotFound(func() string {
		return "Page not found!"
	})

	http.Handle("/", m)
}

func age(birthday time.Time) int {
	now := time.Now()
	years := now.Year() - birthday.Year()
	if now.YearDay() < birthday.YearDay() {
		years--
	}
	return years
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
