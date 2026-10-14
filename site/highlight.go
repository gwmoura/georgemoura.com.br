package main

import (
	"bytes"
	"html"
	"html/template"
	"regexp"
	"strings"
	"sync"

	"github.com/alecthomas/chroma/v2"
	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

//go:generate go run ./cmd/chromacss

// Blocos <pre><code class="linguagem">...</code></pre> dos posts.
var codeBlock = regexp.MustCompile(`(?s)<pre[^>]*>\s*<code([^>]*)>(.*?)</code>\s*</pre>`)
var classAttr = regexp.MustCompile(`class="([^"]*)"`)

var codeFormatter = chromahtml.New(chromahtml.WithClasses(true))

// Cache do HTML dos posts já destacado, por FriendlyId.
var highlightedPosts sync.Map

// highlightCode aplica o destaque de sintaxe do Chroma a cada bloco de código do HTML.
func highlightCode(src string) string {
	return codeBlock.ReplaceAllStringFunc(src, func(block string) string {
		m := codeBlock.FindStringSubmatch(block)
		code := html.UnescapeString(m[2])

		tokens, err := chroma.Coalesce(lexerFor(m[1], code)).Tokenise(nil, code)
		if err != nil {
			return block
		}
		var buf bytes.Buffer
		if err := codeFormatter.Format(&buf, styles.Fallback, tokens); err != nil {
			return block
		}
		return buf.String()
	})
}

// lexerFor escolhe o lexer pela classe do <code> (language-x, lang-x ou x),
// pelo conteúdo ou, por fim, texto puro.
func lexerFor(attrs string, code string) chroma.Lexer {
	if m := classAttr.FindStringSubmatch(attrs); m != nil {
		for _, class := range strings.Fields(m[1]) {
			name := strings.TrimPrefix(strings.TrimPrefix(class, "language-"), "lang-")
			if l := lexers.Get(name); l != nil {
				return l
			}
		}
	}
	if l := lexers.Analyse(code); l != nil {
		return l
	}
	return lexers.Fallback
}

// renderPost renderiza o template do post, aplica o destaque de código e guarda
// o resultado em cache em produção (em desenvolvimento, reflete edições na hora).
func renderPost(t *template.Template, post Post) (template.HTML, error) {
	if c, ok := highlightedPosts.Load(post.FriendlyId); ok {
		return c.(template.HTML), nil
	}

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "posts/"+post.FriendlyId, post); err != nil {
		return "", err
	}
	content := template.HTML(highlightCode(buf.String()))
	if env() == "production" {
		highlightedPosts.Store(post.FriendlyId, content)
	}
	return content, nil
}
