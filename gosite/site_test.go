package site

import (
        "testing"
        "net/http"
        "net/http/httptest"
        "time"
        "strings"
        // "fmt"
)


func DoRequest(method string, path string) *httptest.ResponseRecorder{
    request, _ := http.NewRequest(method, path, nil)
    response := httptest.NewRecorder()

    http.DefaultServeMux.ServeHTTP(response, request)
    return response
}

func TestSite(t *testing.T){
    response := DoRequest("GET","/")
    if response.Code != http.StatusOK {
        t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "200", response.Code)
    }

    response = DoRequest("GET","/feed")
    if response.Code != http.StatusOK{
        t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "200", response.Code)
    }

    response = DoRequest("GET","/blog")
    if response.Code != http.StatusOK{
        t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "200", response.Code)
    }

    response = DoRequest("GET","/contatos")
    if response.Code != http.StatusOK{
        t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "200", response.Code)
    }

    response = DoRequest("GET","/curriculo")
    if response.Code != http.StatusOK{
        t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "200", response.Code)
    }

    //test to view a post
    response = DoRequest("GET",posts[2].Link)
    if response.Code != http.StatusOK{
        t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "200", response.Code)
    }

    if !strings.Contains(response.Body.String(), posts[2].Title){
        t.Fatalf("Response body did not contain expected %v:\n\t", posts[2].Title)
    }

    //test page not found
    response = DoRequest("GET","/xyz")
    if response.Code != http.StatusNotFound{
        t.Fatalf("Response body did not contain expected %v:\n\tbody: %v", "404", response.Code)
    }
}

func TestAge(t *testing.T) {
        years := age(time.Date(1990, 5, 14, 0, 0, 0, 0, time.UTC))

        if years != 25 {
                t.Fatalf("Year different from %d", 25)
        }
}

func TestGetFeed(t *testing.T) {
    feed := getFeed()
    if !strings.Contains(feed, "Programador Web e Mobile "){
        t.Fatalf("Response body not contais %v", "Programador Web e Mobile ")
    }
}

func TestGetPosts(t *testing.T){
    allPosts := getPosts()
    totalposts := len(allPosts)
    if allPosts[totalposts-1].Title != "Uma boa IDE para Ruby on Rails"{
        t.Fatalf("Post Title is diferrent. expected: %v\n actual: %v", "Uma boa IDE para Ruby on Rails", allPosts[totalposts-1].Title)
    }
}

func TestGetPostByName(t *testing.T){
    p := getPostByName("chrome-os-no-meu-pc")
    if p.Title!="Chrome OS no meu pc"{
        t.Fatalf("Post Title different. expected: %v\n actual", "Chrome OS no meu pc", p.Title)
    }
}
