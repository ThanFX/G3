package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	homeTemplate, err := template.New("index.html").Delims("{{%", "%}}").ParseFiles("./public/index.html")

	//homeTemplate, err := template.ParseFiles("./public/index.html")
	if err != nil {
		log.Fatal("Ошибка парсинга шаблона index.html: ", err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTemplate.ExecuteTemplate(w, "index", nil)
}
