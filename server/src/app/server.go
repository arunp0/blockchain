package main

import (
	"app/blockchain"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	bc := blockchain.BlockChain{}
	bc.CreateBlock(1, "0")

	http.Handle("/mine_block", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get previous block
		pBlock := bc.PreviousBlock()
		// get proof of work (mining)
		proof := bc.ProofOfWork(pBlock.GetProof())
		// add the block to block chain
		b := bc.CreateBlock(proof, pBlock.Hash())

		log.Println("Successfully Added to blockchain")

		json.NewEncoder(w).Encode(b)
	}))

	http.Handle("/get_chain", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(bc.GetChain())
	}))

	http.Handle("/is_valid", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		valid := bc.IsChainValid()
		if valid {
			fmt.Fprintf(w, "Block chain is valid")
		} else {
			fmt.Fprintf(w, "Block chain is messed up")
		}
	}))

	log.Println("Now server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}
