package blockchain

import(
	"time" 
	"github.com/aman-zulfiqar/BookBlockChain/models"
)

type Blockchain struct {
	Blocks []*Block
}

func CreateBlock(prevBlock *Block, checkoutItem models.BookCheckout) *Block {
	block := &Block{}
	block.Pos = prevBlock.Pos + 1
	block.Timestamp = time.Now().String()
	block.Data = checkoutItem
	block.PrevHash = prevBlock.Hash
	block.generateHash()
	return block
}

func GenesisBlock() *Block {
	return CreateBlock(&Block{}, models.BookCheckout{IsGenesis: true})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

func (bc *Blockchain) AddBlock(data models.BookCheckout) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	block := CreateBlock(prevBlock, data)
	if validBlock(block, prevBlock) {
		bc.Blocks = append(bc.Blocks, block)
	}
}

func validBlock(block, prevBlock *Block) bool {
	if prevBlock.Hash != block.PrevHash {
		return false
	}
	if !block.validateHash() {
		return false
	}
	if prevBlock.Pos+1 != block.Pos {
		return false
	}
	return true
}
