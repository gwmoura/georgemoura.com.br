package site

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

type PageMetaTag struct {
	Description string
	Keywords    string
}

var metaTags = map[string]PageMetaTag{
	"uma-boa-ide-para-ruby-on-rails": {
		Description: "Meus caros estou escrevendo esse post para falar sobre uma dificuldade que tive logo quando resolvi programar para ruby &#8211; uma IDE. Comecei meus primeiros estudos utilizando o Eclipse e fazendo os teste em meu próprio notebook, mas assim que comecei a testar meu códigos em servidores me deparei com problemas: A versão do rails",
		Keywords:    "ide,ruby,rails,sublimetext,eclipse,netbeans",
	},
	"como-deletar-registros-duplicados-no-banco-de-dados-deixando-apenas-um": {
		Description: "Pessoal, vim deixar um post bem pequeno, mas que foi minha salvação e achei legal compartilhar. Estou em um projeto e acabei precisando fazer uma migração de dados, o problema é que meu script acabou duplicando as informações no banco de dados, imagine o problema, a sorte é que era um banco de desenvolvimento e",
		Keywords:    "mysql,banco de dados",
	},
	"chrome-os-no-meu-pc": {
		Description: "Atenção faça os testes usando uma máquina virtual ou um pc que não tenha utilidade Aí galera venho falar para vocês sobre o Chrome OS, o sistema operacional da Google, acabei de instalar ele no meu pen-drive e já estou utilizando, a propósito estou escrevendo este post nele. Eu estava empolgado para comprar um chromebook,",
		Keywords:    "chrome os,chrome,virtualbox",
	},
	"iniciando-com-o-docker": {
		Description: "O que é o Docker? &#8220;An open platform for distributed applications&#8221; -&#160;Uma plataforma aberta para aplicativos distribuídos. Essa é a definição no site, mas não fica muito clara no início, porém faz sentido depois. O Docker é uma ferramenta para distribuir aplicações em containers de forma rápida e leve, ou seja você não precisa mais",
		Keywords:    "docker,nodejs",
	},
	"criando-vms-na-google-cloud-platform": {
		Description: "Neste tutorial vou mostrar como criar VMs na nuvem da Google (Google Cloud Platform). O primeiro passo é acessar: cloud.google.com e clicar em &#8220;Sign in&#8221; no canto superior direito. Após o login (com sua conta google) clique em &#8220;My Console&#8221; no canto superior direito. Ao clicar nesse link você irá para a página de projetos.",
		Keywords:    "google,cloud,vm,google cloud platform,compute engine",
	},
}

func init() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/", func() string {
		r.HTML(200, "home")
	})

	m.Get("/blog/", func() string {
		r.HTML(200, "blog", metaTags)
	})

	m.Get("/:postname", func(params martini.Params) string {
		postname := params["postname"]
		r.HTML(200, "posts/"+postname, metaTags[postname].Description, metaTags[postname].Keywords)
	})

	http.Handle("/", m)
}
