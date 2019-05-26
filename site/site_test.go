package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func DoRequest(method string, path string, f func(martini.Params, render.Render)) *httptest.ResponseRecorder {
	// req, _ := http.NewRequest(method, path, nil)
	// rr := httptest.NewRecorder()
	// m := martini.New()
	// MyInstance.ServeHTTP(rr, req)

	// http.DefaultServeMux.ServeHTTP(rr, req)

	// handler := http.HandlerFunc(f)
	// handler.ServeHTTP(rr, req)

	// return rr

	rr := httptest.NewRecorder()
	// m := martini.New()
	setUp()
	setRoutes()
	MyInstance.ServeHTTP(rr, (*http.Request)(nil))

	return rr
}

func TestSite(t *testing.T) {
	response := DoRequest("GET", "/", index)
	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}

	response = DoRequest("GET", "/feed", pageFeed)
	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}

	response = DoRequest("GET", "/blog", blog)
	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}

	response = DoRequest("GET", "/contatos", contatos)
	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}

	response = DoRequest("GET", "/curriculo", showCVPage)
	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}

	response = DoRequest("GET", "/curriculo/en", showCVPage)
	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}

	response = DoRequest("GET", "/cv", showCVPage)
	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}

	response = DoRequest("GET", "/cv/br", showCVPage)
	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}

	// test if all posts not have a broken link and test show for post
	for _, post := range getPosts() {
		response = DoRequest("GET", post.Link, getPost)
		if response.Code != http.StatusOK {
			t.Fatalf("Response code did not contain expected %v:\n\tCode: %v \t URL: %v", "200", response.Code, post.Link)
		}
		body := response.Body.String()
		if !strings.Contains(body, post.Title) {
			t.Fatalf("Response body did not contain expected %v , body %v:\n\t", post.Title, body)
		}
	}

	//test page not found
	response = DoRequest("GET", "/xyz", getPost)
	if response.Code != http.StatusNotFound {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "404", response.Code)
	}
}

func TestAge(t *testing.T) {
	myage := 29
	years := age(time.Date(1990, 5, 14, 0, 0, 0, 0, time.UTC))

	if years != myage {
		t.Fatalf("Year different from %d", myage)
	}
}

func TestGetFeed(t *testing.T) {
	feed := getFeed()
	if !strings.Contains(feed, "Programador Web e Mobile ") {
		t.Fatalf("Response body not contais %v", "Programador Web e Mobile ")
	}
}

func TestGetPosts(t *testing.T) {
	allPosts := getPosts()
	totalposts := len(allPosts)
	if allPosts[totalposts-1].Title != "Uma boa IDE para Ruby on Rails" {
		t.Fatalf("Post Title is diferrent. expected: %v\n actual: %v", "Uma boa IDE para Ruby on Rails", allPosts[totalposts-1].Title)
	}
}

func TestGetPostByName(t *testing.T) {
	p := getPostByName("chrome-os-no-meu-pc")
	if p.Title != "Chrome OS no meu pc" {
		t.Fatalf("Post Title different. expected: %v\n actual: %v", "Chrome OS no meu pc", p.Title)
	}
}
