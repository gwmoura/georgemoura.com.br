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

	response = DoRequest("GET", "/euli")
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
	birthday := time.Date(1990, 5, 14, 0, 0, 0, 0, time.UTC)
	expected := ageAt(birthday, time.Now())
	years := age(birthday)

	if years != expected {
		t.Fatalf("Age expected %d but receive %d", expected, years)
	}
}

func TestAgeAt(t *testing.T) {
	birthday := time.Date(1990, 5, 14, 0, 0, 0, 0, time.UTC)
	cases := []struct {
		now      time.Time
		expected int
	}{
		{time.Date(2025, 5, 13, 0, 0, 0, 0, time.UTC), 34},
		{time.Date(2025, 5, 14, 0, 0, 0, 0, time.UTC), 35},
		{time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC), 35},
		{time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC), 35},
		// leap year: May 13 has the same YearDay as May 14 in a common year
		{time.Date(2028, 5, 13, 0, 0, 0, 0, time.UTC), 37},
		{time.Date(2028, 5, 14, 0, 0, 0, 0, time.UTC), 38},
	}

	for _, c := range cases {
		years := ageAt(birthday, c.now)
		if years != c.expected {
			t.Fatalf("Age at %v expected %d but receive %d", c.now.Format("2006-01-02"), c.expected, years)
		}
	}
}

func TestBooks(t *testing.T) {
	response := DoRequest("GET", "/euli")
	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}

	for _, book := range books {
		path := "/euli/" + book.FriendlyId
		response = DoRequest("GET", path)
		if response.Code != http.StatusOK {
			t.Fatalf("Response code did not contain expected %v:\n\tCode: %v \t URL: %v", "200", response.Code, path)
		}
		body := response.Body.String()
		if !strings.Contains(body, book.Title) {
			t.Fatalf("Response body did not contain expected %v , URL %v", book.Title, path)
		}
	}

	response = DoRequest("GET", "/euli/xyz")
	if response.Code != http.StatusNotFound {
		t.Fatalf("Response code did not contain expected %v:\n\tfound: %v", "404", response.Code)
	}
}

func TestNatal(t *testing.T) {
	response := DoRequest("GET", "/natal")
	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tbody: %v", "200", response.Code)
	}
}

func TestLetsencryptNotFound(t *testing.T) {
	response := DoRequest("GET", "/.well-known/acme-challenge/xyz")
	if response.Code != http.StatusNotFound {
		t.Fatalf("Response code did not contain expected %v:\n\tfound: %v", "404", response.Code)
	}
}

func TestGetPostsSkipPages(t *testing.T) {
	for _, post := range getPosts() {
		if post.PostType == "page" {
			t.Fatalf("getPosts should not return pages, found: %v", post.FriendlyId)
		}
	}
}

func TestGetPostByNameNotFound(t *testing.T) {
	p := getPostByName("xyz")
	if p.FriendlyId != "" {
		t.Fatalf("Expected empty post but receive %v", p.FriendlyId)
	}
}

func TestGetBookByNameNotFound(t *testing.T) {
	p := getBookByName("xyz")
	if p.FriendlyId != "" {
		t.Fatalf("Expected empty book but receive %v", p.FriendlyId)
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

func TestGetBookByName(t *testing.T) {
	p := getBookByName("implantacao-de-help-desk-e-service-desk")
	if p.Title != "Implantação de Help Desk e Service Desk" {
		t.Fatalf("Post Title different. expected: %v\n actual: %v", "Chrome OS no meu pc", p.Title)
	}
}
