package blockchain

import (
	"encoding/json"
	"log"
	"net/http"
)

// type blockchain implements ServeHTTP to become Handler.
func (b *blockchain) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		w.Write([]byte("GET request is called\n"))
		b.GetBlockchain(w, r)
	}

	if r.Method == http.MethodPost {
		w.Write([]byte("Post request is called\n"))

		b.AddBlock(w, r)
	}
}

func (b *blockchain) AddBlock(w http.ResponseWriter, r *http.Request) {
	// the argument to decode must be non-nil pointer.
	blk := &block{}
	err := json.NewDecoder(r.Body).Decode(blk)

	if err != nil {
		panic(err)
	}

	b.Blocks = append(b.Blocks, *blk)
}

func (b *blockchain) GetBlockchain(w http.ResponseWriter, r *http.Request) {
	bc := b.GetBlocks()
	err := json.NewEncoder(w).Encode(bc)

	if err != nil {
		log.Fatal(err)
	}
}
