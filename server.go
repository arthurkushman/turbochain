package main

import (
	"github.com/arthurkushman/turbochain/handlers"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/recoilme/pudge"
)

func run() error {
	handlerService := handlers.GetHandlerService()
	muxRouter := mux.NewRouter()
	// base handlers
	muxRouter.HandleFunc("/", handlerService.HandleWriteBlock).Methods("POST")
	muxRouter.HandleFunc("/one", handlerService.HandleGetBlock).Queries("index", "{index:[0-9]+}", "hash", "{hash}").Methods("GET")
	muxRouter.HandleFunc("/last", handlerService.HandleGetLastBlocks).Methods("GET")

	httpAddr := os.Getenv("ADDR")
	log.Println("Listening on ", os.Getenv("ADDR"))
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        muxRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func main() {
	defer pudge.CloseAll()
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(run())
}
