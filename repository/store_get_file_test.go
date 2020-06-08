package repository

import (
	"github.com/arthurkushman/turbochain/components"
	"github.com/recoilme/pudge"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestStoreGetLast(t *testing.T) {
	defer pudge.CloseAll()
	err := os.Setenv("BLOCK_CHAIN_DB", "test")
	assert.NoError(t, err)
	defer os.Remove("test")
	defer os.Remove("test.idx")

	title := "foo bar"
	desc := "foo bar baz"
	block := components.Block{
		Index:     1,
		Timestamp: "2020-06-03 23:00:00 +0000 UTC m=+0.000000000",
		Hash:      "",
		PrevHash:  "",
		Msg: &components.Message{
			BPM:         123,
			Title:       &title,
			Description: &desc,
		},
	}
	newBlock, err := components.GenerateBlock(block, components.Message{
		BPM:         12345,
		Title:       &title,
		Description: &desc,
	})
	assert.NoError(t, err)
	NewStoreService().Store(newBlock)

	got := NewGetService().Get(int64(newBlock.Index), newBlock.Hash)
	assert.Equal(t, got.Index, newBlock.Index)
	assert.Equal(t, got.Hash, newBlock.Hash)

	list := NewGetService().GetLast(1)
	assert.Greater(t, len(list), 0)
}

func TestGetGenesisBlock(t *testing.T) {
	gen := GetGenesisBlock()
	assert.Equal(t, gen.Index, 0)
	assert.Equal(t, gen.PrevHash, "")
}
