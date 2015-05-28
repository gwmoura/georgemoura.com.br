package feed

import (
	"fmt"
	"github.com/gorilla/feeds"
	"net/http"
	"time"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	feed := &feeds.Feed{
		Title:       "George Moura Blog",
		Link:        &feeds.Link{Href: "http://georgemoura.com.br"},
		Description: "Programador PHP, Ruby, Java, Javascript e Mobile",
		Author:      &feeds.Author{"George Moura", "george@georgemoura.com.br"},
		Created:     now,
	}

	feed.Items = []*feeds.Item{
		&feeds.Item{
			Title:       "Limiting Concurrency in Go",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go/"},
			Description: "A discussion on controlled parallelism in golang",
			Author:      &feeds.Author{"Jason Moiron", "jmoiron@jmoiron.net"},
			Created:     now,
		},
		&feeds.Item{
			Title:       "Logic-less Template Redux",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/logicless-template-redux/"},
			Description: "More thoughts on logicless templates",
			Created:     now,
		},
		&feeds.Item{
			Title:       "Idiomatic Code Reuse in Go",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/idiomatic-code-reuse-in-go/"},
			Description: "How to use interfaces <em>effectively</em>",
			Created:     now,
		},
	}
	rss, err := feed.ToRss()
	if err != nil {
		fmt.Fprint(w, "Error!")
	}

	fmt.Fprint(w, rss)
}
