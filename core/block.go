package core

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Block single holder
type Block struct {
	Index       int
	PrevHash    string
	Date        time.Time
	Data        interface{}
	CurrentHash string
}

// Blocks type of slice Block
type Blocks []*Block

// NewBlock initialize Block Pointer/ constructor
func NewBlock(index int, data interface{}, date time.Time) *Block {
	return &Block{
		Index: index,
		Date:  date,
		Data:  data,
	}
}

// createHash create hash to every blocks
func (b *Block) createHash() string {
	d := fmt.Sprintf("%v%v%v%v", b.Index, b.PrevHash, b.Date, b.Data)
	h := sha256.New()
	h.Write([]byte(d))
	return fmt.Sprintf("%x", h.Sum(nil))
}
