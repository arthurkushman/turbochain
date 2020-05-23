package repository

import (
	"github.com/arthurkushman/turbochain/components"
	"github.com/recoilme/pudge"
	"os"
	"time"
)

// Getter defines getter methods for storage repository
type Getter interface {
	Get(hash string) components.Block
	GetLast(n int) []components.Block
}

// GetService receiver type for select methods
type GetService struct{}

// NewGetService creates get service
func NewGetService() *GetService {
	return &GetService{}
}

// Get expose block by index and hash
func (s *GetService) Get(index int64, hash string) *components.Block {
	var block components.Block
	pudge.Get(os.Getenv("BLOCK_CHAIN_DB"), index, &block)
	if block.Hash == hash {
		return &block
	}

	return nil
}

// GetLast gets last n blocks
func (s *GetService) GetLast(n int) []components.Block {
	db := os.Getenv("BLOCK_CHAIN_DB")
	keys, _ := pudge.Keys(db, nil, n, 0, false)

	blocks := make([]components.Block, 0)
	for _, key := range keys {
		var block components.Block
		pudge.Get(db, key, &block)
		blocks = append(blocks, block)
	}

	return blocks
}

// GetGenesisBlock creates and returns genesis Block of chain
func GetGenesisBlock() components.Block {
	t := time.Now()
	return components.Block{
		0,
		t.String(),
		0,
		components.CalculateHash(
			components.Block{
				0,
				t.String(),
				0,
				"",
				""}),
		""}
}
