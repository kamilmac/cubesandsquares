package main

import (
	// "fmt"
	"github.com/boltdb/bolt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"path"
)

const (
	DB_PATH string = "cubes.db"
	gDriveURL string = "https://33af57eeb9498ac053d3e355288e591ca01d5a7b.googledrive.com/host/0B-f7h8x-3DuZfktDNmtfMzlEN0hqdk1ORzJzRTNWQXlsNndmVEZTVGpjWUJkb1FGRDJGcE0/"
	PASSWORD string = "123"
)

var db *bolt.DB

func main() {
	db = openDb(DB_PATH)
	defer db.Close()
	
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	router.GET("/", indexHandler)
	router.GET("/print/:printId", printHandler)
	router.GET("/admin", adminHandler)
	router.POST("/addprint", addPrintHandler)

	log.Fatal(http.ListenAndServe(":3000", router))
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := map[string]interface{}{
		"prints": getAllPrints(),
		"gdriveurl": gDriveURL,
	}
	
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl := serveContent("index.html")
	tmpl.ExecuteTemplate(w, "layout", data)
}

func printHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data := map[string]interface{}{
		"prints": getPrint(string(ps.ByName("printId"))),
	}
	
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl := serveContent("print.html")
	tmpl.ExecuteTemplate(w, "layout", data)
}

func adminHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := map[string]interface{}{
		"prints": getAllPrints(),
	}
	
	tmpl := serveContent("admin.html")
	tmpl.ExecuteTemplate(w, "layout", data)
}

func addPrintHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()
	file := r.FormValue("file")
	title := r.FormValue("title")
	p := Print{"", file, title}
	p.savePrint("123")
}

func serveContent(file string) (tmpl *template.Template) {
	layout := path.Join("templates", "layout.html")
	body := path.Join("templates", file)

	tmpl, _ = template.ParseFiles(layout, body)
	return
}
