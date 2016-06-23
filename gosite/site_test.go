package site

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	// "fmt"
)

func DoRequest(method string, path string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, path, nil)
	response := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(response, request)
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

	response = DoRequest("GET", "/cv")
	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}

	// test if all posts not have a broken link and test show for post
	for _, post := range getPosts() {
		response = DoRequest("GET", post.Link)
		if response.Code != http.StatusOK {
			t.Fatalf("Response code did not contain expected %v:\n\tCode: %v \t URL: %v", "200", response.Code, post.Link)
		}
		if !strings.Contains(response.Body.String(), post.Title) {
			t.Fatalf("Response body did not contain expected %v:\n\t", post.Title)
		}
	}

	//test page not found
	response = DoRequest("GET", "/xyz")
	if response.Code != http.StatusNotFound {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "404", response.Code)
	}
}

func TestAge(t *testing.T) {
	myage := 26
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
