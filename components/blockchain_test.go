package components

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateHash(t *testing.T) {
	title := "foo bar"
	desc := "foo bar baz"
	hash := CalculateHash(Block{
		Index:     1,
		Timestamp: "2020-06-03 23:00:00 +0000 UTC m=+0.000000000",
		Hash:      "",
		PrevHash:  "",
		Msg: &Message{
			BPM:         123,
			Title:       &title,
			Description: &desc,
		},
	})
	assert.Equal(t, hash, "a8681a5a61a25e48bd1fd830384b2e2f41965fecb73732b0f9b76c163e0320e5")
}

func TestGenerateBlock(t *testing.T) {
	title := "foo bar"
	desc := "foo bar baz"
	block := Block{
		Index:     1,
		Timestamp: "2020-06-03 23:00:00 +0000 UTC m=+0.000000000",
		Hash:      "",
		PrevHash:  "",
		Msg: &Message{
			BPM:         123,
			Title:       &title,
			Description: &desc,
		},
	}
	newBlock, err := GenerateBlock(block, Message{
		BPM:         12345,
		Title:       &title,
		Description: &desc,
	})
	assert.NoError(t, err)
	assert.Equal(t, block.Index+1, newBlock.Index)
	assert.Equal(t, block.Hash, newBlock.PrevHash)
}

func TestIsBlockValid(t *testing.T) {
	title := "foo bar"
	desc := "foo bar baz"
	block := Block{
		Index:     1,
		Timestamp: "2020-06-03 23:00:00 +0000 UTC m=+0.000000000",
		Hash:      "",
		PrevHash:  "",
		Msg: &Message{
			BPM:         123,
			Title:       &title,
			Description: &desc,
		},
	}
	newBlock, err := GenerateBlock(block, Message{
		BPM:         12345,
		Title:       &title,
		Description: &desc,
	})
	assert.NoError(t, err)

	isValid := IsBlockValid(newBlock, block)
	assert.True(t, isValid)
}
