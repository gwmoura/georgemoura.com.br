<?php
require 'vendor/autoload.php';

$app = new \Slim\Slim();

// Google App Engine doesn't set $_SERVER['PATH_INFO']
$app->environment['PATH_INFO'] = $_SERVER['REQUEST_URI'];


if (isset($_SERVER['SERVER_SOFTWARE']) && strpos($_SERVER['SERVER_SOFTWARE'],'Google App Engine') !== false) {
  $debug = false;
  $mode = 'production';
  $log_enable = false;
  $log_level = \Slim\Log::INFO;
}
else
{
  $debug = true;
  $mode = 'development';
  $log_enable = false;
  $log_level = \Slim\Log::DEBUG;
}

$app->config(array(
  'debug' 		=> $debug,
  'mode' 		=> $mode,
  'templates.path' 	=> '../templates/',
  'log.enable' 		=> $log_enable,
  'log.level' 		=> $log_level,
  'view' 		=> '\Slim\LayoutView',
  'layout' 		=> '../templates/layouts/main.php'
));

$app->get('/', function() use ($app) {
  $app->etag('home');
  $app->expires(604800);
  
  $meta_description = "Blog de George Moura";
  $meta_keywords = "";

  $data = array(
    'meta_description' => $meta_description,
    'meta_keywords' => $meta_keywords
  );

  $app->render('home.php', $data);
});

$app->get('/blog/', function() use ($app) {
  $app->etag('blog');
  $app->expires(2592000);

  $data = array(
    'meta_description' => "George Moura Blog",
    'meta_keywords' => "blog,ide,google cloud,ruby,python,go,docker,php,sql,mysql"
  );

  $app->render('blog.php', $data);
});

$app->get('/:postname/', function($postname) use ($app) {

  $app->etag($postname);
  $app->expires(2592000);

  $posts = scandir('../templates/posts', 1);

  $data = array(
    'uma-boa-ide-para-ruby-on-rails' => array(
      'meta_description' => "Meus caros estou escrevendo esse post para falar sobre uma dificuldade que tive logo quando resolvi programar para ruby &#8211; uma IDE. Comecei meus primeiros estudos utilizando o Eclipse e fazendo os teste em meu próprio notebook, mas assim que comecei a testar meu códigos em servidores me deparei com problemas: A versão do rails", 
      'meta_keywords' => "ide,ruby,rails,sublimetext,eclipse,netbeans"
    ),
    'como-deletar-registros-duplicados-no-banco-de-dados-deixando-apenas-um' => array(
      'meta_description' => "Pessoal, vim deixar um post bem pequeno, mas que foi minha salvação e achei legal compartilhar. Estou em um projeto e acabei precisando fazer uma migração de dados, o problema é que meu script acabou duplicando as informações no banco de dados, imagine o problema, a sorte é que era um banco de desenvolvimento e", 
      'meta_keywords' => "mysql,banco de dados"
    ),
    'chrome-os-no-meu-pc' => array(
      'meta_description' => "Atenção faça os testes usando uma máquina virtual ou um pc que não tenha utilidade Aí galera venho falar para vocês sobre o Chrome OS, o sistema operacional da Google, acabei de instalar ele no meu pen-drive e já estou utilizando, a propósito estou escrevendo este post nele. Eu estava empolgado para comprar um chromebook,", 
      'meta_keywords' => "chrome os,chrome,virtualbox"
    ),
    'iniciando-com-o-docker' => array(
      'meta_description' => "O que é o Docker? &#8220;An open platform for distributed applications&#8221; -&#160;Uma plataforma aberta para aplicativos distribuídos. Essa é a definição no site, mas não fica muito clara no início, porém faz sentido depois. O Docker é uma ferramenta para distribuir aplicações em containers de forma rápida e leve, ou seja você não precisa mais", 
      'meta_keywords' => "docker,nodejs"
    ),
    'criando-vms-na-google-cloud-platform' => array(
      'meta_description' => "Neste tutorial vou mostrar como criar VMs na nuvem da Google (Google Cloud Platform). O primeiro passo é acessar: cloud.google.com e clicar em &#8220;Sign in&#8221; no canto superior direito. Após o login (com sua conta google) clique em &#8220;My Console&#8221; no canto superior direito. Ao clicar nesse link você irá para a página de projetos.", 
      'meta_keywords' => "google,cloud,vm,google cloud platform,compute engine"
    ),
  );
  foreach($posts as $post)
  {
    if(strripos($post,$postname))
    {
      $app->render('posts/'.$post, $data[$postname]);
    }
  }
  
});


$app->run();
