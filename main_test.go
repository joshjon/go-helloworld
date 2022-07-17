package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	go NewServer()
	resp, err := http.DefaultClient.Get("http://localhost:8080")
	if err != nil {
		t.Fatal(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "Hello World!\n" {
		t.FailNow()
	}
}
