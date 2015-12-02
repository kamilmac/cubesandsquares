package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"path"
)

const (
	DB_PATH  string = "cubes.db"
	PASSWORD string = "123"
)

var db *bolt.DB

func main() {
	db = openDb(DB_PATH)
	defer db.Close()
	
	fmt.Println("SAVED PRINTs: ", getAllPrints())

	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	router.GET("/", indexHandler)
	router.GET("/print/:printId", printHandler)
	router.GET("/admin", adminHandler)

	log.Fatal(http.ListenAndServe(":3000", router))
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	varmap := map[string]interface{}{
		"var1": "value",
		"var2": 100,
	}
	
	tmpl := serveContent("index.html")
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.ExecuteTemplate(w, "layout", varmap)
	tmpl.Execute(w, nil)
}

func printHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "Print: ", ps.ByName("printId"))
}

func adminHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Welcome Admin!")
}

func serveContent(file string) (tmpl *template.Template) {
	layout := path.Join("templates", "layout.html")
	body := path.Join("templates", file)

	tmpl, _ = template.ParseFiles(layout, body)
	return
}
