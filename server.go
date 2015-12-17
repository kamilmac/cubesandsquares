package main

import (
	// "fmt"
	"github.com/boltdb/bolt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"path"
	"encoding/json"
	// "io/ioutil"
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
	router.POST("/delprint", delPrintHandler)

	log.Fatal(http.ListenAndServe(":3000", router))
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := map[string]interface{}{
		"prints": getAllPrints(),
		"gdriveurl": gDriveURL,
	}
	tmpl := serveContent("index.html")
	tmpl.ExecuteTemplate(w, "layout", data)
}

func printHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data := map[string]interface{}{
		"prints": getPrint(string(ps.ByName("printId"))),
	}
	tmpl := serveContent("print.html")
	tmpl.ExecuteTemplate(w, "layout", data)
}

func adminHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := map[string]interface{}{
		"prints": getAllPrints(),
		"gdriveurl": gDriveURL,
	}
	tmpl := serveContent("admin.html")
	tmpl.ExecuteTemplate(w, "layout", data)
}

func addPrintHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// body, _ := ioutil.ReadAll(r.Body)
	decoder := json.NewDecoder(r.Body)
	
	var print struct {
		Password string
		Title string
		File string
	}
    decoder.Decode(&print)
	p := Print{"", 0, print.File, print.Title}
	p.savePrint(print.Password)
	
	json.NewEncoder(w).Encode(p)
}

func delPrintHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	success := false
	decoder := json.NewDecoder(r.Body)
	
	var print struct {
		Password string
		Id string
	}
	
    decoder.Decode(&print)
	success = deletePrint(print.Id, print.Password)
	
	response := struct {
		Success bool
	}{
		success,
	}
	
	json.NewEncoder(w).Encode(response)
}

func serveContent(file string) (tmpl *template.Template) {
	layout := path.Join("templates", "layout.html")
	body := path.Join("templates", file)

	tmpl, _ = template.ParseFiles(layout, body)
	return
}
