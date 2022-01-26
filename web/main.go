package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/emrahsariboz/blockchain-restapi/web/blockchain"
)

func main() {

	port := flag.String("port", ":9090", "Port number of server")

	flag.Parse()

	blockchain := blockchain.NewBlockchain()

	// Create Multiplexer (router)
	mux := http.NewServeMux()

	// Add Handler to end-point
	mux.Handle("/", blockchain)

	log.Printf("Server started at port %s", *port)

	// Start Server
	http.ListenAndServe(*port, mux)

}
