package main

import (
	"net/http"
	"strings"
	"testing"
)

func TestHighlightCodeKnownLanguage(t *testing.T) {
	out := highlightCode(`<p>texto</p><pre>
<code class="sql">select * from users;</code>
</pre>`)

	if !strings.HasPrefix(out, "<p>texto</p>") {
		t.Errorf("HTML fora do bloco de código foi alterado: %q", out)
	}
	if !strings.Contains(out, `<pre class="chroma">`) {
		t.Errorf("esperava <pre class=\"chroma\">, veio %q", out)
	}
	if !strings.Contains(out, `<span class="k">select</span>`) {
		t.Errorf("esperava a palavra-chave SQL destacada, veio %q", out)
	}
}

func TestHighlightCodeLanguagePrefixes(t *testing.T) {
	for _, class := range []string{"language-go", "lang-go", "go"} {
		out := highlightCode(`<pre><code class="` + class + `">package main</code></pre>`)
		if !strings.Contains(out, `<span class="kn">package</span>`) {
			t.Errorf("classe %q: esperava destaque de Go, veio %q", class, out)
		}
	}
}

func TestHighlightCodeWithoutLanguage(t *testing.T) {
	out := highlightCode(`<pre><code>algum texto qualquer</code></pre>`)

	if !strings.Contains(out, `<pre class="chroma">`) {
		t.Errorf("esperava bloco do Chroma, veio %q", out)
	}
	if !strings.Contains(out, "algum texto qualquer") {
		t.Errorf("texto do código não foi preservado: %q", out)
	}
}

func TestHighlightCodeEscapesSpecialChars(t *testing.T) {
	out := highlightCode(`<pre><code class="unknown-lang">if a &lt; b &amp;&amp; c &gt; d</code></pre>`)

	if strings.Contains(out, "a < b") || strings.Contains(out, "c > d") {
		t.Errorf("caracteres especiais não foram escapados: %q", out)
	}
	for _, s := range []string{"&lt;", "&amp;&amp;", "&gt;"} {
		if !strings.Contains(out, s) {
			t.Errorf("esperava %q no resultado, veio %q", s, out)
		}
	}
	if strings.Contains(out, "&amp;lt;") {
		t.Errorf("escape duplicado: %q", out)
	}
}

func TestHighlightCodeKeepsPlainPre(t *testing.T) {
	in := `<pre>sem code</pre>`
	if out := highlightCode(in); out != in {
		t.Errorf("<pre> sem <code> não deveria mudar, veio %q", out)
	}
}

func TestPostWithHighlightedCode(t *testing.T) {
	response := DoRequest("GET", "/explicando-mysql-explain")
	if response.Code != http.StatusOK {
		t.Fatalf("esperava 200, veio %v", response.Code)
	}
	body := response.Body.String()
	if !strings.Contains(body, `<pre class="chroma">`) {
		t.Error("post deveria ter blocos de código do Chroma")
	}
	if strings.Contains(body, "highlight.js") || strings.Contains(body, "hljs") {
		t.Error("post não deveria mais carregar o highlight.js")
	}
	if !strings.Contains(body, "/css/chroma.css") {
		t.Error("layout deveria carregar /css/chroma.css")
	}
}
