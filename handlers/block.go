package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aman-zulfiqar/BookBlockChain/blockchain"
	"github.com/aman-zulfiqar/BookBlockChain/models"
)

var BlockChain *blockchain.Blockchain

func InitBlockchain() {
	BlockChain = blockchain.NewBlockchain()
}

func WriteBlock(w http.ResponseWriter, r *http.Request) {
	var checkoutItem models.BookCheckout
	if err := json.NewDecoder(r.Body).Decode(&checkoutItem); err != nil {
		http.Error(w, "could not write block", http.StatusInternalServerError)
		log.Println("Decode error:", err)
		return
	}

	BlockChain.AddBlock(checkoutItem)
	resp, _ := json.MarshalIndent(checkoutItem, "", " ")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetBlockchain(w http.ResponseWriter, r *http.Request) {
	jbytes, err := json.MarshalIndent(BlockChain.Blocks, "", " ")
	if err != nil {
		http.Error(w, "could not get blockchain", http.StatusInternalServerError)
		return
	}
	w.Write(jbytes)
}