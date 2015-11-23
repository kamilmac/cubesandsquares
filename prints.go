package main

import (
    "fmt"
    "encoding/json"
)

type (
    Print struct {
        Id []byte       "json:'id'"   
        File []byte     "json:'file'"
        Title []byte    "json:'title'"
    }
    Prints []Print
)

func (p *Print) savePrint(pass string) {
    if pass == PASSWORD {
        p.Id = getUid()
        json, _ := json.Marshal(p)
        put([]byte("prints"), p.Id, json)
    } else {
        fmt.Println("Invalid password.")
    }
} 

func getAllPrints() (allPrints Prints) {
    list := getAll([]byte("prints"))
    for i:= 0; i < len(list); i++ {
        p := Print{}
        json.Unmarshal(list[i], p)
        allPrints = append(allPrints, p)
        return nil
    }
    return
}