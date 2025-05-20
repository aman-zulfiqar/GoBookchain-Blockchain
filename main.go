package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
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

func writeblock(w *http.ResponseWriter, r *http.Request){
	var checkoutitem BookCheckout

	if err := json.NewDecoder(r.Body).Decode(&BookCheckout); err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not Write the book checkout:%v", err)
		w.Write([]byte("could not Write the book"))
		return
	}

	BlockChain.AddBlock(checkoutitem)

	resp, err := json.MarshalIndent(checkoutitem, "", " ")
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not marshal payload:%v", err)
		w.Write([]byte("could not save the checkoutdata"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func newBook(w *http.ResponseWriter, r *http.Request) {
	var book Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not create new book:%v", err)
		w.Write([]byte("could not create the new book"))
		return
	}

	h := md5.New()
	io.WriteString(h, book.ISBN + book.PublishDate)
	book.ID = fmt.Sprintf("%x", h.Sum(nil))

	resp, err := json.MarshalIndent(book, "", " ")
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not marshal the oayload:%v", err)
		w.Write([]byte("could not save the book data"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getblockchain).Methods("GET")
	r.HandleFunc("/", writeblock).Methods("POST")
	r.HandleFunc("/new", newblock).Methods("POST")

	log.Println("The sere ver is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
