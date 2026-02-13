# Instruções do GitHub Copilot - georgemoura.com.br

## Visão Geral do Projeto

Este é um site pessoal e blog desenvolvido por George Moura. O projeto é uma aplicação web escrita em **Go** que roda no **Google App Engine**, utilizando o framework **Martini** para roteamento e renderização de templates.

## Stack Tecnológica

### Backend
- **Linguagem**: Go 1.12+
- **Runtime**: Google App Engine (Go 1.15)
- **Framework Web**: Martini (`github.com/go-martini/martini`)
- **Renderização**: `martini-contrib/render` com templates HTML
- **Compressão**: `martini-contrib/gzip`
- **RSS Feed**: `github.com/gorilla/feeds`

### Frontend
- **Framework CSS**: Materialize CSS
- **JavaScript**: Vanilla JS
- **Fontes**: Material Design Icons, Roboto

### Infraestrutura
- **Hospedagem**: Google App Engine
- **Storage**: Google Cloud Storage (para imagens/assets)
- **CI/CD**: Codeship

## Estrutura do Projeto

```
/site/                    # Código principal da aplicação
  ├── site.go            # Ponto de entrada, rotas e handlers
  ├── database.go        # Array de posts (dados estáticos)
  ├── models.go          # Structs (Post)
  ├── app.yaml           # Configuração do App Engine
  ├── cron.yaml          # Configuração de cron jobs
  ├── go.mod             # Dependências Go
  ├── templates/         # Templates HTML
  │   ├── layout.tmpl   # Layout principal
  │   ├── home.tmpl     # Página inicial
  │   ├── blog.tmpl     # Lista de posts
  │   ├── posts/        # Posts em HTML
  │   └── cv/           # Currículos (PT/EN)
  ├── public/            # Assets estáticos
  │   ├── css/          # Estilos
  │   ├── js/           # Scripts
  │   ├── images/       # Imagens
  │   └── font/         # Fontes
  └── letsencrypt/       # Certificados SSL
/scripts/                # Scripts de automação
  ├── deploy            # Script de deploy
  ├── serve.sh          # Servidor de desenvolvimento
  └── run_tests.sh      # Execução de testes
/images/                 # Imagens para upload no GCS
```

## Padrões de Código

### Adicionando Novos Posts

1. **Criar arquivo HTML** em `/site/templates/posts/{nome-do-post}.html`
2. **Adicionar entrada** no array `posts` em `/site/database.go`:

```go
{
    Title:       "Título do Post",
    Description: "Descrição que aparece na meta tag e RSS",
    Keywords:    "palavra1,palavra2,palavra3",
    Link:        "/nome-do-post/",
    CreatedAt:   time.Date(YYYY, time.Month, DD, HH, MM, SS, 0, time.UTC).Local(),
    FriendlyId:  "nome-do-post",
},
```

3. Posts são ordenados por data de criação (mais recente primeiro)
4. O FriendlyId deve corresponder ao nome do arquivo HTML sem extensão

### Convenções de Nomenclatura

- **URLs**: kebab-case com trailing slash (ex: `/nome-do-post/`)
- **FriendlyId**: kebab-case sem slash (ex: `nome-do-post`)
- **Arquivos de Posts**: kebab-case.html (ex: `nome-do-post.html`)
- **Templates**: snake_case ou kebab-case com extensão `.tmpl` ou `.html`

### Template Functions

O projeto define funções customizadas disponíveis nos templates:

- `assetsHost`: Retorna URL do Google Cloud Storage
- `host`: Retorna domínio do site
- `env`: Retorna ambiente (production/development)
- `inProduction`: Retorna boolean se está em produção
- `gaAccount`: Retorna ID do Google Analytics (apenas em produção)
- `hasString`: Verifica se string contém substring

Exemplo de uso em templates:
```html
<link href="{{assetsHost}}/css/style.css" rel="stylesheet">
{{if inProduction}}
  <!-- Google Analytics -->
{{end}}
```

### Estrutura de Posts HTML

Posts devem conter:
- Tags semânticas HTML5
- Links com `target="_blank"` para URLs externas
- Imagens hospedadas no Google Cloud Storage
- Code blocks com syntax highlighting quando aplicável

## Rotas Principais

```
GET /                    # Página inicial
GET /blog                # Lista de todos os posts
GET /:postname           # Post individual
GET /curriculo/:locale   # Currículo (pt-br ou en)
GET /cv/:locale          # Alias para currículo
GET /feed                # RSS Feed
GET /contatos            # Página de contatos
GET /euli                # Página de livros
GET /euli/:bookname      # Livro específico
```

## Comandos Importantes

```bash
# Desenvolvimento
make serve               # Inicia servidor local (porta 8080)
make start               # Alias para make serve

# Testes
make tests               # Executa testes Go

# Deploy
make deploy-images       # Upload de imagens para GCS
make deploy              # Deploy completo (imagens + código)
```

## Regras de Deploy

- **App Engine**: Deploy via `gcloud app deploy`
- **Imagens**: São armazenadas no Google Cloud Storage
- **Cache**: Headers de cache configurados para 1 hora (max-age=3600)
- **Expiração padrão**: 30 minutos (app.yaml)
- **HTTPS**: Sempre forçado (secure: always no app.yaml)
- **Scaling**: Máximo 1 instância, 15 requisições concorrentes

## Variáveis de Ambiente

- `ENV`: Define ambiente (`production` ou desenvolvimento)
- `PORT`: Porta do servidor (default: 8080)

## Práticas de Desenvolvimento

### Ao Modificar Templates
- Sempre testar localmente com `make serve`
- Sempre parar o servidor e reiniciar para refletir mudanças
- Verificar responsividade (Materialize CSS)
- Validar HTML semântico
- Testar meta tags para SEO

### Ao Adicionar Dependências Go
- Atualizar `go.mod`
- Executar `go mod tidy`
- Verificar compatibilidade com App Engine Go 1.15

### Ao Modificar Rotas
- Atualizar em `setRoutes()` no `site.go`
- Criar handler correspondente
- Adicionar testes se necessário

## Testes

- Arquivo de testes: `/site/site_test.go`
- Executar com: `make tests` ou `./scripts/run_tests.sh`
- Testes devem cobrir handlers principais e funções auxiliares

## Assets Estáticos

- **CSS**: Materialize CSS + style.css customizado
- **JavaScript**: Materialize.js + site.js customizado
- **CDN**: Google Cloud Storage para otimização
- **Fontes**: Material Design Icons e Roboto hospedadas localmente

## SEO e Meta Tags

Cada página/post deve definir:
- `Title`: Título da página
- `Description`: Meta description
- `Keywords`: Palavras-chave (separadas por vírgula)
- Estrutura de dados para RSS feed

## Considerações Especiais

1. **Posts Legados**: Projeto migrado de WordPress, mantém compatibilidade com URLs antigas
2. **Multilíngue**: Currículos disponíveis em PT e EN
3. **Let's Encrypt**: Diretório `.well-known/acme-challenge` para renovação SSL
4. **Error Handling**: Páginas customizadas para erros e over quota

## Skills Úteis para o Projeto

O GitHub Copilot pode ser instruído para executar tarefas específicas através de comandos naturais. Aqui estão algumas skills úteis para este projeto:

### 1. Criar Novo Post
**Comando**: `@workspace crie um novo post sobre [tópico]`

O Copilot deve:
- Criar arquivo HTML em `/site/templates/posts/`
- Adicionar entrada no array `posts` em `/site/database.go`
- Usar convenções de nomenclatura kebab-case
- Gerar estrutura HTML semântica com meta tags
- Incluir data atual e keywords relevantes

### 2. Validar Estrutura de Post
**Comando**: `@workspace valide o post [nome-do-post]`

O Copilot deve verificar:
- Arquivo HTML existe em `/site/templates/posts/`
- Entrada correspondente existe em `database.go`
- FriendlyId corresponde ao nome do arquivo
- Meta tags (Title, Description, Keywords) estão presentes
- Links externos têm `target="_blank"`

### 3. Otimizar SEO
**Comando**: `@workspace analise SEO do post [nome-do-post]`

O Copilot deve avaliar:
- Presença de meta tags completas
- Tamanho ideal de Title (50-60 caracteres)
- Tamanho ideal de Description (150-160 caracteres)
- Uso adequado de headings (H1, H2, H3)
- Alt text em imagens
- Keywords relevantes

### 4. Gerar Template de Post
**Comando**: `@workspace gere template de post em branco`

Criar estrutura básica:
```html
<section>
  <header>
    <h1>[Título]</h1>
    <time datetime="YYYY-MM-DD">[Data]</time>
  </header>
  <article>
    <!-- Conteúdo aqui -->
  </article>
</section>
```

### 5. Atualizar Currículo
**Comando**: `@workspace adicione nova experiência ao currículo`

O Copilot deve:
- Identificar arquivos `/site/templates/cv/br.tmpl` e `en.tmpl`
- Manter formatação consistente
- Atualizar ambas versões (PT e EN)
- Preservar estrutura do Materialize CSS

### 6. Verificar Compatibilidade de Deploy
**Comando**: `@workspace verifique compatibilidade para deploy`

Checklist automático:
- `go.mod` está atualizado
- Testes em `site_test.go` passam
- `app.yaml` configurado corretamente
- Variável `ENV` definida
- Assets referenciando `assetsHost` corretamente

### 7. Listar Posts por Período
**Comando**: `@workspace liste posts de [ano/mês]`

O Copilot deve:
- Analisar `database.go`
- Filtrar posts por data (`CreatedAt`)
- Retornar lista organizada cronologicamente

### 8. Migrar Post do WordPress
**Comando**: `@workspace migre post do wordpress [conteúdo]`

O Copilot deve:
- Converter HTML do WordPress para estrutura do projeto
- Ajustar URLs de imagens para Google Cloud Storage
- Remover shortcodes e plugins específicos do WP
- Criar entrada em `database.go`
- Manter compatibilidade de URLs

### 9. Gerar Script de Deploy
**Comando**: `@workspace crie script de deploy para [ambiente]`

Automatizar:
- Upload de imagens para GCS
- Build e deploy no App Engine
- Validação pré-deploy
- Rollback se necessário

### 10. Análise de Performance
**Comando**: `@workspace analise performance do site`

Verificar:
- Uso de gzip habilitado
- Cache headers configurados
- Assets servidos do CDN
- Tamanho de imagens
- Minificação de CSS/JS

## Como Pedir ao Copilot

Use comandos naturais começando com `@workspace`:

```
@workspace crie um post sobre "Docker em Produção" com keywords docker, devops, containers
@workspace valide todos os posts criados em 2026
@workspace adicione nova skill JavaScript ao currículo
@workspace otimize imagens da pasta /images/2020/08/
@workspace gere RSS feed atualizado
```

## Referências Úteis

- Documentação Martini: https://github.com/go-martini/martini
- Materialize CSS: http://materializecss.com/
- App Engine Go: https://cloud.google.com/appengine/docs/go

---

**Autor**: George Moura (gwmoura@gmail.com)
**Domínio**: georgemoura.com.br
**Repositório**: github.com/gwmoura/georgemoura.com.br
