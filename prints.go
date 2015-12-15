package main

import (
    "fmt"
    "encoding/json"
    "time"
)

type (
    Print struct {
        Id string           `json:"id"`  
        Timestamp int32     `json:"timestamp"`
        File string         `json:"file"`
        Title string        `json:"title"`
    }
    Prints []Print
)

func (p *Print) savePrint(pass string) bool {
    if valid(pass) {
        p.Id = getUid()
        p.Timestamp = int32(time.Now().Unix())
        j, _ := json.Marshal(p)
        put("prints", p.Id, j)
        return true
    }
    return false
}

func getAllPrints() Prints {
    var allPrints Prints
    
    list := getAll("prints")
    p := Print{}
    for _, v := range list {
        json.Unmarshal([]byte(v), &p)
        allPrints = append(allPrints, p)
    }

    return allPrints
}

func getPrint(id string) (p Print) {
    v := get("prints", id)
    json.Unmarshal(v, &p)
    return
}

func deletePrint(id, pass string) bool {
    if valid(pass) {
        fmt.Println("ID: ", id)
        delete("prints", id)
        return true
    }
    return false
}

func valid(pass string) bool {
    if pass != PASSWORD {
        fmt.Println("Invalid password.")
        return false    
    }
    return true
}