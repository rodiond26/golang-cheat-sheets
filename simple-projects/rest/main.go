package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:Title`
	Desc    string `json:desc`
	Content string `json:content`
}

type Articles []Article

func main() {
	handeRequests()
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Title 1", Desc: "Desc 1", Content: "contetnt 1"},
		Article{Title: "Title 2", Desc: "Desc 2", Content: "contetnt 2"},
		Article{Title: "Title 3", Desc: "Desc 3", Content: "contetnt 3"},
		Article{Title: "Title 4", Desc: "Desc 4", Content: "contetnt 4"},
		Article{Title: "Title 5", Desc: "Desc 5", Content: "contetnt 5"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint ")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Homepage endpoint hit")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Post endpoint hit")
}

func handeRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
