package handlers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/aman-zulfiqar/BookBlockChain/models"
)

func NewBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "could not create new book", http.StatusInternalServerError)
		log.Println("Decode error:", err)
		return
	}

	h := md5.New()
	io.WriteString(h, book.ISBN+book.PublishDate)
	book.ID = fmt.Sprintf("%x", h.Sum(nil))

	resp, _ := json.MarshalIndent(book, "", " ")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}