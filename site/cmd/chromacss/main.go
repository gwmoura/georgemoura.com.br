// Gera public/css/chroma.css com os estilos claro e escuro do destaque de código.
// Uso (em site/): go generate
package main

import (
	"bytes"
	"log"
	"os"
	"regexp"

	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/styles"
)

const (
	lightStyle = "github"
	darkStyle  = "github-dark"
	output     = "public/css/chroma.css"
)

// Cada regra do Chroma vem como "/* Nome */ .seletor { ... }".
var selector = regexp.MustCompile(`(?m)^(/\* [^*]* \*/ )?\.`)

func css(style string, prefix string) []byte {
	s := styles.Get(style)
	if s == styles.Fallback && style != "fallback" {
		log.Fatalf("estilo %q não encontrado", style)
	}
	var buf bytes.Buffer
	if err := html.New(html.WithClasses(true)).WriteCSS(&buf, s); err != nil {
		log.Fatal(err)
	}
	return selector.ReplaceAll(buf.Bytes(), []byte("${1}"+prefix+" ."))
}

func main() {
	var out bytes.Buffer
	out.WriteString("/* Gerado por cmd/chromacss (go generate). Não edite à mão. */\n\n")
	// Os seletores claro e escuro nunca se sobrepõem, para que propriedades de um
	// estilo (ex.: fundo de .err) não vazem para o outro.
	theme := func(name, style, other string) {
		out.WriteString("/* Tema " + name + " escolhido no toggle: " + style + " */\n")
		out.Write(css(style, `[data-theme="`+name+`"]`))
		out.WriteString("\n/* Tema " + name + " pela preferência do sistema, sem escolha no toggle */\n")
		out.WriteString("@media (prefers-color-scheme: " + name + ") {\n")
		out.Write(css(style, `:root:not([data-theme="`+other+`"])`))
		out.WriteString("}\n\n")
	}
	theme("light", lightStyle, "dark")
	theme("dark", darkStyle, "light")

	if err := os.WriteFile(output, out.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}
}
