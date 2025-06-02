package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aman-zulfiqar/BookBlockChain/handlers"
	"github.com/gorilla/mux"
)

func main() {
	handlers.InitBlockchain()

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.GetBlockchain).Methods("GET")
	r.HandleFunc("/", handlers.WriteBlock).Methods("POST")
	r.HandleFunc("/new", handlers.NewBook).Methods("POST")

	go func() {
		for _, block := range handlers.BlockChain.Blocks {
			fmt.Printf("Prev. hash: %x\n", block.PrevHash)
			bytes, _ := json.MarshalIndent(block.Data, "", " ")
			fmt.Printf("Data: %v\n", string(bytes))
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Println()
		}
	}()

	log.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
