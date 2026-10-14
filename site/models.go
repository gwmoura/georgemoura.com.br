package main

import (
	"html/template"
	"time"
)

// PostPage é o post já renderizado, com o código destacado, dentro do layout.
type PostPage struct {
	Post
	Content template.HTML
}

// Post - basic struct for content of site
type Post struct {
	Title       string
	Description string
	Link        string
	Keywords    string
	AuthorName  string
	AuthorEmail string
	CreatedAt   time.Time
	PostType    string
	FriendlyId  string
}
