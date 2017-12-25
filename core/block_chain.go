package core

import (
	"time"
)

type Chain interface {
	AddBlock(block *Block)
	GetLastBlock() *Block
	GetBlocks() Blocks
	IsChainStillValid() bool
}

// BlockChain model
type BlockChain struct {
	Blocks Blocks
}

// NewBlockChain initialize BlockChain Pointer/ constructor
func NewBlockChain() Chain {
	//date created, we use January this year
	date, _ := time.Parse(time.RFC3339, "2017-01-01T22:08:41+00:00")
	//create genesis block
	genesisBlock := NewBlock(0, "Genesis Block", date)
	genesisBlock.PrevHash = "0"
	genesisBlock.CurrentHash = genesisBlock.createHash()

	var blocks Blocks
	blocks = append(blocks, genesisBlock)

	return &BlockChain{
		Blocks: blocks,
	}
}

// AddBlock function, to add Block to BlockChain's Blocks
func (b *BlockChain) AddBlock(block *Block) {
	// set new block's PrevHash with latest block's current hash
	block.PrevHash = b.GetLastBlock().CurrentHash

	// set current hash from created hash
	block.CurrentHash = block.createHash()
	b.Blocks = append(b.Blocks, block)
}

// GetLastBlock function, get latest index of BlockChain's Blocks
func (b *BlockChain) GetLastBlock() *Block {
	return b.Blocks[len(b.Blocks)-1]
}

// GetBlocks function, return Blocks
func (b *BlockChain) GetBlocks() Blocks {
	return b.Blocks
}

// IsChainStillValid function that return error, will return false if one of block's value changed
func (b *BlockChain) IsChainStillValid() bool {
	for i := 1; i < len(b.Blocks); i++ {
		currentBlock := b.Blocks[i]
		prevBlock := b.Blocks[i-1]

		if currentBlock.CurrentHash != currentBlock.createHash() {
			return false
		}

		if currentBlock.PrevHash != prevBlock.CurrentHash {
			return false
		}
	}

	return true
}
