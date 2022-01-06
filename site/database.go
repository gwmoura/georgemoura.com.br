package main

import (
	"time"
)

var posts = []Post{
	{
		Title:       "Desenvolvedor Web e Mobile",
		Description: "Blog de George Moura",
		Keywords:    "",
		PostType:    "page",
		FriendlyId:  "home",
	},
	{
		Title:       "Currículo",
		Description: "Currículo George Moura",
		Keywords:    "",
		PostType:    "page",
		FriendlyId:  "cv",
	},
	{
		Description: "O MySQL Explain é uma ferramenta usada explicar o plano de execução que o SGBD do MySQL usa para executar uma query. Com essa ferramenta é possível entender quais índices o banco utiliza na query, quais os relacionamentos usados, quantas linhas foram examidas para retornar os dados, se foi ou não necessário ler todos os dados da tabela.",
		Keywords:    "mysql,database,bancode de daods,mysql explain",
		Title:       "Explicando o MySQL Explain",
		Link:        "/explicando-mysql-explain/",
		CreatedAt:   time.Date(2020, time.December, 25, 17, 07, 17, 0, time.UTC).Local(),
		FriendlyId:  "explicando-mysql-explain",
	},
	{
		Description: "Vim é um editor de texto muito poderoso",
		Keywords:    "vim,editor de textos",
		Title:       "Começando com Vim",
		Link:        "/comecando-com-vim/",
		CreatedAt:   time.Date(2020, time.August, 23, 19, 10, 0, 0, time.UTC).Local(),
		FriendlyId:  "comecando-com-vim",
	},
	{
		Description: "Gosto muito de ser preguiçoso e sempre que me vejo fazendo alguma coisa repetidas vezes fico imaginando formas de automatizar. Isso me ajuda a ser mais produtivo.",
		Keywords:    "scripts,automatização,bash,shell,ambiente de desenvolvimento",
		Title:       "Automatizando seu ambiente de desenvolvimento",
		Link:        "/automatizando-seu-ambiente-de-desenvolvimento/",
		CreatedAt:   time.Date(2020, time.August, 5, 11, 20, 39, 0, time.UTC).Local(),
		FriendlyId:  "automatizando-seu-ambiente-de-desenvolvimento",
	},
	{
		Description: "NGINX é um servidor web que também pode ser usado como proxy reverso e balanceador de carga, proxy de e-mail e servidor de cache. O Nginx é conhecido por ser um servidor de alta performance, estável, de fácil configuração e baixo consumo de recursos",
		Keywords:    "nginx,web server,load balancer,reverse proxy",
		Title:       "NGINX - O que é e como usar",
		Link:        "/nginx-o-que-e-e-como-usar/",
		CreatedAt:   time.Date(2020, time.July, 19, 0, 9, 10, 0, time.UTC).Local(),
		FriendlyId:  "nginx-o-que-e-e-como-usar",
	},
	{
		Description: "Micro serviços ou Monolíticos? O que é melhor? A resposta para essa pergunta é a mesma para muitas para muitas perguntas do universo de desenvolvimento de software: Depende :)",
		Keywords:    "microservices,monolithics,micro-servicos,monolitos,arquitetura",
		Title:       "Micro serviços ou monolitos? Eis a questão",
		Link:        "/micro-servicos-ou-monolitos-eis-a-questao/",
		CreatedAt:   time.Date(2020, time.July, 12, 23, 35, 00, 0, time.UTC).Local(),
		FriendlyId:  "micro-servicos-ou-monolitos-eis-a-questao",
	},
	{
		Description: "Neste post vou falar um pouco sobre Docker Compose e como usá-lo em desenvolvimento",
		Keywords:    "php,docker,docker-compose,desenvolvimento",
		Title:       "Desenvolvendo com Docker Compose PHP",
		Link:        "/desenvolvendo-com-docker-compose-php/",
		CreatedAt:   time.Date(2016, time.April, 4, 13, 05, 0, 0, time.UTC).Local(),
		FriendlyId:  "desenvolvendo-com-docker-compose-php",
	},
	{
		Description: "Neste post vou falar um pouco sobre Docker Compose e como usá-lo em desenvolvimento",
		Keywords:    "ruby,docker,docker-compose,desenvolvimento",
		Title:       "Desenvolvendo com Docker Compose Ruby",
		Link:        "/desenvolvendo-com-docker-compose-ruby/",
		CreatedAt:   time.Date(2016, time.April, 4, 13, 05, 0, 0, time.UTC).Local(),
		FriendlyId:  "desenvolvendo-com-docker-compose-ruby",
	},
	{
		Description: "Ruby 2.3.0 foi lançado recentemente e eu resolvi entender um pouco sobre benchmark e fazer um comparando as versões 2.2.3 e a 2.3.0",
		Keywords:    "ruby, benchmark",
		Title:       "Ruby Benchmark 2.2.3 VS 2.3.0",
		Link:        "/ruby-benchmark-223-vs-230/",
		CreatedAt:   time.Date(2015, time.December, 26, 14, 22, 0, 0, time.UTC).Local(),
		FriendlyId:  "ruby-benchmark-223-vs-230",
	},
	{
		Description: "Estou escrevendo este post para não esquecer como configurar novas resoluções de monitores no meu ubuntu.",
		Keywords:    "ubuntu, monitor",
		Title:       "Adicionando resolução do monitor no ubuntu 14.04",
		Link:        "/adicionando-resolucao-do-monitor-no-ubuntu-1404/",
		CreatedAt:   time.Date(2015, time.December, 15, 13, 46, 0, 0, time.UTC).Local(),
		FriendlyId:  "adicionando-resolucao-do-monitor-no-ubuntu-1404",
	},
	{
		Description: "Neste tutorial irei mostrar como instalar o docker no ubuntu 14.04",
		Keywords:    "docker,ubuntu",
		Title:       "Instalando Docker no Ubuntu 14.04",
		Link:        "/instalando-docker-no-ubuntu-1404/",
		CreatedAt:   time.Date(2015, time.November, 22, 20, 15, 0, 0, time.UTC).Local(),
		FriendlyId:  "instalando-docker-no-ubuntu-1404",
	},
	{
		Description: "Neste tutorial vou mostrar como criar VMs na nuvem da Google (Google Cloud Platform). O primeiro passo é acessar: cloud.google.com e clicar em \"Sign in\" no canto superior direito. Após o login (com sua conta google) clique em \"My Console\" no canto superior direito. Ao clicar nesse link você irá para a página de projetos.",
		Keywords:    "google,cloud,vm,google cloud platform,compute engine",
		Title:       "Criando VMs na Google Cloud Platform",
		Link:        "/criando-vms-na-google-cloud-platform/",
		CreatedAt:   time.Date(2015, time.March, 17, 21, 2, 37, 0, time.UTC).Local(),
		FriendlyId:  "criando-vms-na-google-cloud-platform",
	},
	{
		Description: "O que é o Docker? \"An open platform for distributed applications\" - Uma plataforma aberta para aplicativos distribuídos. Essa é a definição no site, mas não fica muito clara no início, porém faz sentido depois. O Docker é uma ferramenta para distribuir aplicações em containers de forma rápida e leve, ou seja você não precisa mais",
		Keywords:    "docker,nodejs",
		Title:       "Iniciando com o Docker",
		Link:        "/iniciando-com-o-docker/",
		CreatedAt:   time.Date(2014, time.October, 13, 22, 6, 47, 0, time.UTC).Local(),
		FriendlyId:  "iniciando-com-o-docker",
	},
	{
		Description: "Atenção faça os testes usando uma máquina virtual ou um pc que não tenha utilidade Aí galera venho falar para vocês sobre o Chrome OS, o sistema operacional da Google, acabei de instalar ele no meu pen-drive e já estou utilizando, a propósito estou escrevendo este post nele. Eu estava empolgado para comprar um chromebook,",
		Keywords:    "chrome os,chrome,virtualbox",
		Title:       "Chrome OS no meu pc",
		Link:        "/chrome-os-no-meu-pc/",
		CreatedAt:   time.Date(2013, time.October, 12, 18, 53, 22, 0, time.UTC).Local(),
		FriendlyId:  "chrome-os-no-meu-pc",
	},
	{
		Description: "Pessoal, vim deixar um post bem pequeno, mas que foi minha salvação e achei legal compartilhar. Estou em um projeto e acabei precisando fazer uma migração de dados, o problema é que meu script acabou duplicando as informações no banco de dados, imagine o problema, a sorte é que era um banco de desenvolvimento e",
		Keywords:    "mysql,banco de dados",
		Title:       "Como deletar registros duplicados no banco de dados, deixando apenas um",
		Link:        "/como-deletar-registros-duplicados-no-banco-de-dados-deixando-apenas-um/",
		CreatedAt:   time.Date(2013, time.October, 5, 12, 20, 19, 0, time.UTC).Local(),
		FriendlyId:  "como-deletar-registros-duplicados-no-banco-de-dados-deixando-apenas-um",
	},
	{
		Description: "Meus caros estou escrevendo esse post para falar sobre uma dificuldade que tive logo quando resolvi programar para ruby &#8211; uma IDE. Comecei meus primeiros estudos utilizando o Eclipse e fazendo os teste em meu próprio notebook, mas assim que comecei a testar meu códigos em servidores me deparei com problemas: A versão do rails",
		Keywords:    "ide,ruby,rails,sublimetext,eclipse,netbeans",
		Title:       "Uma boa IDE para Ruby on Rails",
		Link:        "/uma-boa-ide-para-ruby-on-rails/",
		CreatedAt:   time.Date(2013, time.July, 4, 0, 23, 47, 0, time.UTC).Local(),
		FriendlyId:  "uma-boa-ide-para-ruby-on-rails",
	},
}

var books = []Post{
	{
		Title:       "Implantação de Help Desk e Service Desk",
		Description: "",
		Keywords:    "",
		PostType:    "book",
		FriendlyId:  "implantacao-de-help-desk-e-service-desk",
		Link:        "/implantacao-de-help-desk-e-service-desk/",
		CreatedAt:   time.Date(2021, time.December, 18, 21, 24, 0, 0, time.UTC).Local(),
	},
	{
		Title:       "Os segredos da mente milionária",
		Description: "",
		Keywords:    "",
		PostType:    "book",
		FriendlyId:  "os-segredos-da-mente-milionaria",
		Link:        "/os-segredos-da-mente-milionaria/",
		CreatedAt:   time.Date(2022, time.January, 5, 21, 14, 0, 0, time.UTC).Local(),
	},
}
