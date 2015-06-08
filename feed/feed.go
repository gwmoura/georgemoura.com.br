package feed

import (
	"fmt"
	"github.com/gorilla/feeds"
	"net/http"
	"time"
)

var authorName string = "George Moura"
var authorEmail string = "george@georgemoura.com.br"

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	feed := &feeds.Feed{
		Title:       "George Moura Blog",
		Link:        &feeds.Link{Href: "http://georgemoura.com.br"},
		Description: "Programador Web e Mobile",
		Author:      &feeds.Author{authorName, authorEmail},
		Created:     now,
	}

	feed.Items = []*feeds.Item{
		&feeds.Item{
			Title:       "Criando VMs na Google Cloud Platform",
			Link:        &feeds.Link{Href: "http://georgemoura.com.br/criando-vms-na-google-cloud-platform/"},
			Description: "<![CDATA[<p>Neste tutorial vou mostrar como criar VMs na nuvem da Google (Google Cloud Platform). O primeiro passo é acessar: cloud.google.com e clicar em &#8220;Sign in&#8221; no canto superior direito. Após o login (com sua conta google) clique em &#8220;My Console&#8221; no canto superior direito. Ao clicar nesse link você irá para a página de projetos. [&#8230;]</p> <p>The post <a rel=\"nofollow\" href=\"http://georgemoura.com.br/criando-vms-na-google-cloud-platform/\">Criando VMs na Google Cloud Platform</a> appeared first on <a rel=\"nofollow\" href=\"http://georgemoura.com.br\">George Moura</a>.</p>]]>",
			Author:      &feeds.Author{authorName, authorEmail},
			Created:     time.Date(2015, time.March, 17, 18, 2, 37, 0, time.UTC).Local(),
		},
		&feeds.Item{
			Title:       "Iniciando com o Docker",
			Link:        &feeds.Link{Href: "http://georgemoura.com.br/iniciando-com-o-docker/"},
			Description: "<![CDATA[<p>O que é o Docker? &#8220;An open platform for distributed applications&#8221; -&#160;Uma plataforma aberta para aplicativos distribuídos. Essa é a definição no site, mas não fica muito clara no início, porém faz sentido depois. O Docker é uma ferramenta para distribuir aplicações em containers de forma rápida e leve, ou seja você não precisa mais [&#8230;]</p> <p>The post <a rel=\"nofollow\" href=\"http://georgemoura.com.br/iniciando-com-o-docker/\">Iniciando com o Docker</a> appeared first on <a rel=\"nofollow\" href=\"http://georgemoura.com.br\">George Moura</a>.</p>]]>",
			Author:      &feeds.Author{authorName, authorEmail},
			Created:     time.Date(2014, time.October, 13, 19, 6, 47, 0, time.UTC).Local(),
		},
		&feeds.Item{
			Title:       "Chrome OS no meu pc",
			Link:        &feeds.Link{Href: "http://georgemoura.com.br/chrome-os-no-meu-pc/"},
			Description: "<![CDATA[<p>Atenção faça os testes usando uma máquina virtual ou um pc que não tenha utilidade Aí galera venho falar para vocês sobre o Chrome OS, o sistema operacional da Google, acabei de instalar ele no meu pen-drive e já estou utilizando, a propósito estou escrevendo este post nele. Eu estava empolgado para comprar um chromebook, [&#8230;]</p> <p>The post <a rel=\"nofollow\" href=\"http://georgemoura.com.br/chrome-os-no-meu-pc/\">Chrome OS no meu pc</a> appeared first on <a rel=\"nofollow\" href=\"http://georgemoura.com.br\">George Moura</a>.</p>]]>",
			Author:      &feeds.Author{authorName, authorEmail},
			Created:     time.Date(2013, time.October, 12, 15, 53, 22, 0, time.UTC).Local(),
		},

		&feeds.Item{
			Title:       "Como deletar registros duplicados no banco de dados, deixando apenas um",
			Link:        &feeds.Link{Href: "http://georgemoura.com.br/como-deletar-registros-duplicados-no-banco-de-dados-deixando-apenas-um/"},
			Description: "<![CDATA[<p>Pessoal, vim deixar um post bem pequeno, mas que foi minha salvação e achei legal compartilhar. Estou em um projeto e acabei precisando fazer uma migração de dados, o problema é que meu script acabou duplicando as informações no banco de dados, imagine o problema, a sorte é que era um banco de desenvolvimento e [&#8230;]</p> <p>The post <a rel=\"nofollow\" href=\"http://georgemoura.com.br/como-deletar-registros-duplicados-no-banco-de-dados-deixando-apenas-um/\">Como deletar registros duplicados no banco de dados, deixando apenas um</a> appeared first on <a rel=\"nofollow\" href=\"http://georgemoura.com.br\">George Moura</a>.</p>]]>",
			Author:      &feeds.Author{authorName, authorEmail},
			Created:     time.Date(2013, time.October, 5, 9, 20, 19, 0, time.UTC).Local(),
		},
		&feeds.Item{
			Title:       "Uma boa IDE para Ruby on Rails",
			Link:        &feeds.Link{Href: "http://georgemoura.com.br/uma-boa-ide-para-ruby-on-rails/"},
			Description: "<![CDATA[<p>Meus caros estou escrevendo esse post para falar sobre uma dificuldade que tive logo quando resolvi programar para ruby &#8211; uma IDE. Comecei meus primeiros estudos utilizando o Eclipse e fazendo os teste em meu próprio notebook, mas assim que comecei a testar meu códigos em servidores me deparei com problemas: A versão do rails [&#8230;]</p> <p>The post <a rel=\"nofollow\" href=\"http://georgemoura.com.br/uma-boa-ide-para-ruby-on-rails/\">Uma boa IDE para Ruby on Rails</a> appeared first on <a rel=\"nofollow\" href=\"http://georgemoura.com.br\">George Moura</a>.</p>]]>",
			Author:      &feeds.Author{authorName, authorEmail},
			Created:     time.Date(2013, time.July, 3, 21, 23, 47, 0, time.UTC).Local(),
		},
	}
	rss, err := feed.ToRss()
	if err != nil {
		fmt.Fprint(w, "Errors!")
	}

	fmt.Fprint(w, rss)
}
