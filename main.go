package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			log.Print("Got a GET, serve /")
			var t = template.Must(template.ParseFiles("template.html"))
			t.Execute(w, nil)
		case "POST":
			log.Print("Got a POST, send a mail and redirect")
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		default:
			errString := fmt.Sprintf("Method %v not allowed", r.Method)
			log.Print(errString)
			http.Error(w, errString, http.StatusMethodNotAllowed)
		}
	})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
