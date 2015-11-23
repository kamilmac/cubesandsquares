package main

import (
    "fmt"
    "log"
    "path"
    "net/http"
    "html/template"
    "time"
    "strings"
    "github.com/julienschmidt/httprouter"
	"github.com/boltdb/bolt"
    "github.com/satori/go.uuid"

    // ADD BOLTDB
)

// WRITE BOLTDB HANDLER
// put, get, getAll, delete, openDb

// Write uuid handler

// Create dummy prints obj

type (
    print struct {
        Id string       "json:'id'"   
        File string     "json:'file'"
        Title string    "json:'title'"
    }
)
const DB_PATH string = "cubes.db"
var db *bolt.DB



func getUid() (id []byte) {
    return uuid.NewV4().Bytes()
}

func openDb(path string) (DB *bolt.DB) {
	DB, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	return
}

func put(bucket, id, value []byte) {
    err := db.Update(func(tx *bolt.Tx) error {
        b, err := tx.CreateBucketIfNotExists(bucket)
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        }
        b.Put([]byte(id), value)
        return nil
    })
    if err != nil {
		log.Fatal(err)
	}
}

func getAll(bucket []byte) (response string){
	result := []string{}
	db.View(func(tx *bolt.Tx) error {
	    b := tx.Bucket([]byte(bucket))
	    b.ForEach(func(k, v []byte) error {
			result = append(result, fmt.Sprintf("\"%s\": %s", k, v))
	        return nil
	    })
	    return nil
	})
	response = fmt.Sprintf("{%s}", (strings.Join(result, ",")))
	return
}





func main() {
    db = openDb(DB_PATH)
    defer db.Close()
    
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