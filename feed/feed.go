package hello

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func getPost(string postname) {
	files, _ := ioutil.ReadDir("../posts/")
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
