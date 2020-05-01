package main

import (
	"fmt"
	"html"
	"log"
	"time"

	//"log"
	"net/http"
)

type book struct {
	id          int    `json:"id"`
	title       string `json:"title"`
	author      string `json:"author"`
	year        string `json:"year"`
	publication string `json:"publication"`
}

type rack struct {
	rackid int     `json:"rackid"`
	books  []*book `json:[]book`
}

func routine1() {
	fmt.Print("routine one")
}

func routinenumbers() {
	for i := 1; i < 10; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Print(i)
	}
}

func routinealphabets() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c", i)
		time.Sleep(450 * time.Millisecond)
	}
}

func main() {
	fmt.Println("hello world")
	fmt.Println("This is project one")
	go routinenumbers()
	go routine1()
	go routinealphabets()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, %q", html.EscapeString(request.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":4567", nil))
}
