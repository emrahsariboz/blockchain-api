package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/emrahsariboz/blockchain-restapi/web/blockchain"
	"github.com/gorilla/mux"
)

func main() {

	port := flag.String("port", ":9090", "Port number of server")

	flag.Parse()

	blockchain := blockchain.NewBlockchain()

	// Create Multiplexer (router)
	//mux := http.NewServeMux()
	sm := mux.NewRouter()

	// Add Handler to end-point
	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", blockchain.GetBlockchain)

	// PUT Handler to end-point
	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/add", blockchain.AddBlock)

	log.Printf("Server started at port %s", *port)

	// Start Server
	http.ListenAndServe(*port, sm)

}
