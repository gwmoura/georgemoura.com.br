package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func DoRequest(method string, path string) *httptest.ResponseRecorder {
	setUp()
	setRoutes()
	req, _ := http.NewRequest(method, path, nil)
	response := httptest.NewRecorder()
	MyInstance.ServeHTTP(response, req)
	return response
}

func TestSite(t *testing.T) {
	response := DoRequest("GET", "/")
	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}

	response = DoRequest("GET", "/feed")
	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}

	response = DoRequest("GET", "/blog")
	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}

	response = DoRequest("GET", "/contatos")
	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}

	response = DoRequest("GET", "/curriculo")
	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}

	response = DoRequest("GET", "/curriculo/en")
	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}

	response = DoRequest("GET", "/cv")
	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}

	response = DoRequest("GET", "/cv/br")
	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}

	// test if all posts not have a broken link and test show for post
	for _, post := range getPosts() {
		response = DoRequest("GET", post.Link)
		if response.Code != http.StatusOK {
			t.Fatalf("Response code did not contain expected %v:\n\tCode: %v \t URL: %v", "200", response.Code, post.Link)
		}
		body := response.Body.String()
		if !strings.Contains(body, post.Title) {
			t.Fatalf("Response body did not contain expected %v , body %v:\n\t", post.Title, body)
		}
	}

	//test page not found
	response = DoRequest("GET", "/xyz")
	body := response.Body.String()
	if response.Code != http.StatusNotFound {
		t.Fatalf("Response code did not contain expected %v:\n\tfound: %v - body: %v", "404", response.Code, body)
	}
}

func TestAge(t *testing.T) {
	myage := 30
	years := age(time.Date(1990, 5, 14, 0, 0, 0, 0, time.UTC))

	if years != myage {
		t.Fatalf("Age expected %d but receive %d", myage, years)
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
	postTitle := "Uma boa IDE para Ruby on Rails"
	if allPosts[totalposts-1].Title != postTitle {
		t.Fatalf("Post Title is diferrent. expected: %v\n actual: %v", postTitle, allPosts[totalposts-1].Title)
	}
}

func TestGetPostByName(t *testing.T) {
	p := getPostByName("chrome-os-no-meu-pc")
	if p.Title != "Chrome OS no meu pc" {
		t.Fatalf("Post Title different. expected: %v\n actual: %v", "Chrome OS no meu pc", p.Title)
	}
}
