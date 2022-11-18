package main

import (
	"fmt"

	"github.com/quankori/go-consensus/core"
)

func main() {

	bc := core.NewBlockchain()

	bc.AddBlock("Add one user")
	bc.AddBlock("Add two user")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
