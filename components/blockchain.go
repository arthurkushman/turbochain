package components

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Block the repetitive block of chain
type Block struct {
	Index     int      `json:"index"`
	Timestamp string   `json:"timestamp"`
	Hash      string   `json:"hash"`
	PrevHash  string   `json:"prev_hash"`
	Msg       *Message `json:"message"`
}

// Message for BPM in POST
type Message struct {
	BPM         int     `json:"bpm"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

// CalculateHash calculates hash based on Block
func CalculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.Msg.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// GenerateBlock generates new block based on prev
func GenerateBlock(prevBlock Block, msg Message) (Block, error) {
	var newBlock Block
	t := time.Now()

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.PrevHash = prevBlock.Hash

	newBlock.Msg = &Message{}
	newBlock.Msg.BPM = msg.BPM
	if msg.Title != nil {
		newBlock.Msg.Title = msg.Title
	}

	if msg.Description != nil {
		newBlock.Msg.Description = msg.Description
	}
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock, nil
}

// IsBlockValid checks whether this block is valid with prev
func IsBlockValid(newBlock, prevBlock Block) bool {
	if prevBlock.Index+1 != newBlock.Index {
		return false
	}

	if prevBlock.Hash != newBlock.PrevHash {
		return false
	}

	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
