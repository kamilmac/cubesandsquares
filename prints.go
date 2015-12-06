package main

import (
    "fmt"
    "encoding/json"
)

type (
    Print struct {
        Id string       `json:"id"`  
        File string     `json:"file"`
        Title string    `json:"title"`
    }
    Prints []Print
)

func (p Print) savePrint(pass string) string {
    if valid(pass) {
        p.Id = getUid()
        j, _ := json.Marshal(p)
        put("prints", p.Id, j)
    }
    return p.Id
}

func getAllPrints() (allPrints Prints) {
    list := getAll("prints")
    p := Print{}
    for _, v := range list {
        json.Unmarshal([]byte(v), &p)
        allPrints = append(allPrints, p)
    }
    return
}

func getPrint(id string) (p Print) {
    v := get("prints", id)
    json.Unmarshal(v, &p)
    return
}

func deletePrint(id, pass string) {
    if valid(pass) {
        delete("prints", id)
    }
}

func valid(pass string) bool {
    if pass != PASSWORD {
        fmt.Println("Invalid password.")
        return false    
    }
    return true
}