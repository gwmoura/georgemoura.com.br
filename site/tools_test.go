package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFerramentas(t *testing.T) {
	response := DoRequest("GET", "/ferramentas")
	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tfound: %v", "200", response.Code)
	}
	index := response.Body.String()
	for _, tool := range tools {
		if !strings.Contains(index, tool.Link) {
			t.Fatalf("Tools page did not link to %v", tool.Link)
		}
		response = DoRequest("GET", tool.Link)
		if response.Code != http.StatusOK {
			t.Fatalf("Response code did not contain expected %v:\n\tCode: %v \t URL: %v", "200", response.Code, tool.Link)
		}
	}
}

func TestMeuIP(t *testing.T) {
	setUp()
	setRoutes()
	req, _ := http.NewRequest("GET", "/ferramentas/meu-ip", nil)
	req.Header.Set("X-Appengine-User-IP", "203.0.113.7")
	response := httptest.NewRecorder()
	MyInstance.ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Fatalf("Response code did not contain expected %v:\n\tfound: %v", "200", response.Code)
	}
	if !strings.Contains(response.Body.String(), "203.0.113.7") {
		t.Fatalf("Response body did not contain the client IP")
	}
	if cache := response.Header().Get("Cache-Control"); cache != "no-store" {
		t.Fatalf("Cache-Control expected no-store but receive %v", cache)
	}
}

func TestClientIP(t *testing.T) {
	cases := []struct {
		name     string
		headers  map[string]string
		remote   string
		expected string
	}{
		{"app engine", map[string]string{"X-Appengine-User-IP": "203.0.113.7", "X-Forwarded-For": "198.51.100.1"}, "10.0.0.1:1234", "203.0.113.7"},
		{"forwarded for", map[string]string{"X-Forwarded-For": "198.51.100.1, 10.0.0.2"}, "10.0.0.1:1234", "198.51.100.1"},
		{"invalid header", map[string]string{"X-Appengine-User-IP": "lixo", "X-Forwarded-For": "lixo"}, "10.0.0.1:1234", "10.0.0.1"},
		{"remote addr ipv6", nil, "[2001:db8::1]:1234", "2001:db8::1"},
	}

	for _, c := range cases {
		req, _ := http.NewRequest("GET", "/", nil)
		req.RemoteAddr = c.remote
		for k, v := range c.headers {
			req.Header.Set(k, v)
		}
		if ip := clientIP(req); ip != c.expected {
			t.Fatalf("%v: expected %v but receive %v", c.name, c.expected, ip)
		}
	}
}

func TestIPVersion(t *testing.T) {
	cases := map[string]string{"203.0.113.7": "IPv4", "2001:db8::1": "IPv6", "lixo": ""}
	for ip, expected := range cases {
		if v := ipVersion(ip); v != expected {
			t.Fatalf("ipVersion(%v) expected %q but receive %q", ip, expected, v)
		}
	}
}

func TestFerramentaNotFound(t *testing.T) {
	response := DoRequest("GET", "/ferramentas/xyz")
	if response.Code != http.StatusNotFound {
		t.Fatalf("Response code did not contain expected %v:\n\tfound: %v", "404", response.Code)
	}
}

func TestGeradorDeSenhas(t *testing.T) {
	response := DoRequest("GET", "/ferramentas/gerador-de-senhas")
	if !strings.Contains(response.Body.String(), `class="password-generator"`) {
		t.Fatalf("Password generator page did not render the form")
	}
}
