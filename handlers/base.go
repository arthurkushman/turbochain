package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/arthurkushman/turbochain/components"
	"github.com/arthurkushman/turbochain/repository"
	"github.com/davecgh/go-spew/spew"
	"io"
	"net/http"
	"strconv"
)

const LastNBlocks = 100

// HandlerService to operate with handlers
type HandlerService struct {
	GetService   *repository.GetService
	StoreService *repository.StoreService
}

// NewHandlerService creates handler service + generates Genesis block and store it in pudge
func NewHandlerService(getService *repository.GetService, storeService *repository.StoreService) *HandlerService {
	if len(getService.GetLast(1)) == 0 { // if there is no even genesis yet
		go func() {
			genesisBlock := repository.GetGenesisBlock()
			spew.Dump(genesisBlock)
			storeService.Store(genesisBlock)
		}()
	}

	return &HandlerService{getService, storeService}
}

// GetHandlerService creates and returns handlers service
func GetHandlerService() *HandlerService {
	return NewHandlerService(repository.NewGetService(), repository.NewStoreService())
}

// HandleGetBlock handles GET request on get block by hash
func (s *HandlerService) HandleGetLastBlocks(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(s.GetService.GetLast(LastNBlocks), "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

// HandleGetBlock handles GET request on get block by hash
func (s *HandlerService) HandleGetBlock(w http.ResponseWriter, r *http.Request) {
	var idx int64
	var bOut []byte
	var err error

	hash := r.FormValue("hash")
	idx, err = strconv.ParseInt(r.FormValue("index"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bOut, err = json.MarshalIndent(s.GetService.Get(idx, hash), "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	io.WriteString(w, string(bOut))
}

// HandleWriteBlock writes 1 block at a time
func (s *HandlerService) HandleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var m components.Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	prevBlock := s.GetService.GetLast(1)[0]
	newBlock, err := components.GenerateBlock(prevBlock, m.BPM)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, m)
		return
	}
	// validate and store block
	if components.IsBlockValid(newBlock, prevBlock) {
		go s.StoreService.Store(newBlock)
		fmt.Println("last valid blocks prev->next: ")
		// dump 2 last blocks
		spew.Dump([]components.Block{prevBlock, newBlock})
	}

	respondWithJSON(w, http.StatusCreated, newBlock)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
