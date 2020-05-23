package repository

import (
	"github.com/arthurkushman/turbochain/components"
	"log"
	"os"

	"github.com/recoilme/pudge"
)

// Setter defines setter methods for storage repository
type Setter interface {
	Store(newBlock components.Block)
}

// StoreService receiver type for storage methods
type StoreService struct{}

// NewStoreService creates store service
func NewStoreService() *StoreService {
	return &StoreService{}
}

// Store saves new Block in a chain
func (s *StoreService) Store(newBlock components.Block) {
	err := pudge.Set(os.Getenv("BLOCK_CHAIN_DB"), newBlock.Index, newBlock)
	if err != nil {
		log.Println(err)
	}
}
