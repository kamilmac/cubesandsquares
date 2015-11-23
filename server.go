package main

import (
    "fmt"
    "log"
    "path"
    "net/http"
    "html/template"

    "github.com/julienschmidt/httprouter"
    // ADD BOLTDB
)

// WRITE BOLTDB HANDLER
// put, get, getAll, delete, openDb

// Write uuid handler

// Create dummy prints obj

type (
    print struct {
        Id string
        File string
        Title string
    }
)

func main() {
    router := httprouter.New()
    router.ServeFiles("/static/*filepath", http.Dir("static"))
    router.GET("/", indexHandler)
    router.GET("/print/:printId", printHandler)
    router.GET("/admin", adminHandler)

    log.Fatal(http.ListenAndServe(":3000", router))
}


func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    tmpl := serveContent("index.html")
    tmpl.ExecuteTemplate(w, "layout", nil)
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