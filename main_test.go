package main

import (
	"testing"
	"reflect"
	"fmt"
)

func TestUid(t *testing.T) {
	uid := getUid()
	if reflect.TypeOf(uid).String() != "[]uint8" {
		t.Error("GetUid not workinn mannn..")
	}
}

func TestDb(t *testing.T) {
	db = openDb("test.db")
	defer db.Close()

	value := "testing"
	put([]byte("test"), getUid(), []byte(value))
	list := getAll([]byte("test"))
	for _, v := range list {
		if reflect.TypeOf(v).String() != "[]uint8" {
			t.Error("Element was not saved properly")
		}
	}
}

func TestSavePrint(t *testing.T) {
	db = openDb("test.db")
	defer db.Close()
	
	p := Print{[]byte("123"),[]byte("file.png"),[]byte("title")}
	fmt.Println("saved print: ", p)
	
	id := p.savePrint("123")
	
	if string(p.File) != string(getPrint(id).File) {
		t.Error("Save&get print fail")
	}
}