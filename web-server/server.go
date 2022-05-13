package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var goCount int = 0
var pyCount int = 0
var goTimer time.Time
var pyTimer time.Time
var pyStarted bool = false
var goStarted bool = false
var count int = 10000

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world!\n", r.URL.Path[1:])
}

func python(w http.ResponseWriter, r *http.Request) {
	if !pyStarted {
		pyStarted = true
		pyTimer = time.Now()
	}
	pyCount += 1
	if pyCount == count {
		fmt.Printf("Python completed %d requests in:", count)
		fmt.Println(time.Now().Sub(pyTimer))
		pyStarted = false
		pyCount = 0
	}
	fmt.Fprint(w, fmt.Sprintf("%d\n", pyCount), r.URL.Path[1:])
}

func golang(w http.ResponseWriter, r *http.Request) {
	if !goStarted {
		goStarted = true
		goTimer = time.Now()
	}
	goCount += 1
	if goCount == count {
		fmt.Printf("Go completed %d requests in:", count)
		fmt.Println(time.Now().Sub(goTimer))
		goStarted = false
		goCount = 0
	}
	fmt.Fprint(w, fmt.Sprintf("%d\n", goCount), r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/python", python)
	http.HandleFunc("/golang", golang)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
