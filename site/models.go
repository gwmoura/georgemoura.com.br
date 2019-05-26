package main

import "time"

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
