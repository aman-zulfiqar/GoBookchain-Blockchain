package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	

	"github.com/aman-zulfiqar/BookBlockChain/models"
)

type Block struct {
	Pos       int
	Data      models.BookCheckout
	Timestamp string
	Hash      string
	PrevHash  string
}

func (b *Block) generateHash() {
	bytes, _ := json.Marshal(b.Data)
	data := strconv.Itoa(b.Pos) + b.Timestamp + string(bytes) + b.PrevHash
	hash := sha256.New()
	hash.Write([]byte(data))
	b.Hash = hex.EncodeToString(hash.Sum(nil))
}

func (b *Block) validateHash() bool {
	original := b.Hash
	b.generateHash()
	return b.Hash == original
}
