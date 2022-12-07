package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/kangsorang/srcoin/blockchain"
)

const (
	port        string = ":4000"
	templateDir string = "templates/"
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.html"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.html"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	log.Fatal(http.ListenAndServe(port, nil))
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{PageTitle: "home", Blocks: blockchain.GetBlockchain().AllBlock()}
	templates.ExecuteTemplate(rw, "home", data)
}
