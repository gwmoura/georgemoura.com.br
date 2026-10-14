package main

import (
	"net"
	"net/http"
	"strings"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// Tool - ferramenta listada em /ferramentas, cada uma com a sua página
type Tool struct {
	Title       string
	Description string
	Keywords    string
	Link        string
	FriendlyId  string
}

var tools = []Tool{
	{
		Title:       "Meu IP",
		Description: "Descubra o seu endereço IP público, o IP com que você aparece na internet.",
		Keywords:    "ip,meu ip,ip publico,ipv4,ipv6,rede",
		Link:        "/ferramentas/meu-ip",
		FriendlyId:  "meu-ip",
	},
	{
		Title:       "Gerador de senhas",
		Description: "Gere senhas aleatórias escolhendo o tamanho e se terão caracteres especiais. A senha é criada no seu navegador e não é enviada para lugar nenhum.",
		Keywords:    "gerador de senhas,senha,senha forte,senha aleatoria,seguranca",
		Link:        "/ferramentas/gerador-de-senhas",
		FriendlyId:  "gerador-de-senhas",
	},
}

// IPPage - dados da página Meu IP
type IPPage struct {
	Tool
	IP      string
	Version string
}

func ferramentas(params martini.Params, r render.Render) {
	s := struct {
		Title       string
		Description string
		Keywords    string
		Link        string
		Tools       []Tool
	}{"Ferramentas", "Ferramentas que uso no dia a dia", "ferramentas,utilitarios", "/ferramentas", tools}
	r.HTML(200, "ferramentas/index", s)
}

// ferramenta - páginas de ferramentas que não precisam de dados do servidor:
// basta ter a entrada em tools e o template ferramentas/{FriendlyId}
func ferramenta(params martini.Params, r render.Render) {
	tool := getToolByName(params["toolname"])
	if tool.FriendlyId == "" {
		r.Text(404, "Not found")
		return
	}
	r.HTML(200, "ferramentas/"+tool.FriendlyId, tool)
}

func meuIP(res http.ResponseWriter, req *http.Request, r render.Render) {
	// A página muda a cada visitante: não pode ficar no cache público de 1 hora
	res.Header().Set("Cache-Control", "no-store")
	ip := clientIP(req)
	r.HTML(200, "ferramentas/meu-ip", IPPage{Tool: getToolByName("meu-ip"), IP: ip, Version: ipVersion(ip)})
}

// clientIP - IP público de quem fez a requisição. No App Engine o IP real vem
// no X-Appengine-User-IP; fora dele, usa o primeiro do X-Forwarded-For e, por
// último, o endereço da conexão.
func clientIP(req *http.Request) string {
	if ip := strings.TrimSpace(req.Header.Get("X-Appengine-User-IP")); net.ParseIP(ip) != nil {
		return ip
	}
	if forwarded := req.Header.Get("X-Forwarded-For"); forwarded != "" {
		if ip := strings.TrimSpace(strings.Split(forwarded, ",")[0]); net.ParseIP(ip) != nil {
			return ip
		}
	}
	host, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		host = req.RemoteAddr
	}
	return host
}

func ipVersion(ip string) string {
	parsed := net.ParseIP(ip)
	switch {
	case parsed == nil:
		return ""
	case parsed.To4() != nil:
		return "IPv4"
	default:
		return "IPv6"
	}
}

func getToolByName(name string) Tool {
	var t Tool
	for _, tool := range tools {
		if tool.FriendlyId == name {
			t = tool
		}
	}
	return t
}
