package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Block struct {
	Pos       string
	Data      BookCheckout
	TimeStrap string
	Hash      string
	PrevHash  string
}

type BookCheckout struct {
	BookID       string `json:"book_id"`
	User         string `json:"user"`
	CheckoutDate string `json:"checkout_date"`
	IsGenesis    bool   `json:"is_genesis"`
}

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishDate string `json:"publish_date"`
	ISBN        string `json:"isbn"`
}

type BlockChain struct {
	blocks []*Block
}

var Blockchain *BlockChain

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getblockchain).Methods("GET")
	r.HandleFunc("/", writeblock).Methods("POST")
	r.HandleFunc("/new", newblock).Methods("POST")

	log.Println("The sere ver is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
