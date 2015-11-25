package main

import (
    "fmt"
    "encoding/json"
)

type (
    Print struct {
        Id []byte       `json:"id"`  
        File []byte     `json:"file"`
        Title []byte    `json:"title"`
    }
    Prints []Print
)

func (p Print) savePrint(pass string) []byte {
    if pass == PASSWORD {
        p.Id = getUid()
        j, _ := json.Marshal(p)
        put([]byte("prints"), p.Id, j)
    } else {
        fmt.Println("Invalid password.")
    }
    return p.Id
} 

func getAllPrints() (allPrints Prints) {
    list := getAll([]byte("prints"))
    for _, v := range list {
        p := Print{}
        json.Unmarshal(v, p)
        allPrints = append(allPrints, p)
        return nil
    }
    return
}

func getPrint(id []byte) (p Print) {
    v := get([]byte("prints"), id)
    json.Unmarshal(v, &p)
    return
}