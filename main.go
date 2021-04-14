// GoCat project main.go
package main

import (
	"html/template"
	"net/http"
	"strings"
)

var mytemplate *template.Template

func main() {

	mytemplate = template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/", redirectToIndex)
	http.HandleFunc("/gocat", index)
	http.HandleFunc("/gocat/", index)
	http.HandleFunc("/gocat/login", login)
	http.HandleFunc("/gocat/setup", setup)
	http.HandleFunc("/gocat/app", app)
	http.HandleFunc("/gocat/download", download)
	http.HandleFunc("/gocat/logout", logout)
	fs := SetCacheHeader(http.FileServer(http.Dir("static")))

	http.Handle("/gocat/static/", http.StripPrefix("/gocat/static/", fs))
	port := getConfigValue("port", ":2009")
	println("Listening on port: ", port)
	if !strings.Contains(port, ":") {
		port = ":" + port
	}
	err := http.ListenAndServe(port, nil)
	if err != nil {
		println("Error: " + err.Error())
	}

}

func SetCacheHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Frame-Options", "SAMEORIGIN")

		h.ServeHTTP(w, r)
	})
}

func redirectToIndex(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "text/html")
	w.Header().Add("encoding", "UTF-8")
	http.Redirect(w, r, "/gocat", http.StatusPermanentRedirect)
}
