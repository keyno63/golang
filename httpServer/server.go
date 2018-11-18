package httpServer

import (
	"fmt"
	"net/http"
	"text/template"
)

// Page is sample struct.
type Page struct {
	Title string
	Count int
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{"Hello World template file.", 1}
	tmpl, err := template.ParseFiles("layout.html") // ParseFilesを使う
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

// Server run http server.
func Server() {
	http.HandleFunc("/", handler)           // ハンドラを登録してウェブページを表示させる
	http.HandleFunc("/temple", viewHandler) // viewHandler
	http.HandleFunc("/hello", handler)      // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":5091", nil)
}
