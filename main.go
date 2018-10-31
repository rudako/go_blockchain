package main

import (
    "fmt"
    "time"
    "./blockchain"
)

func main() {
    bc := blockchain.NewBlockchain()
    bc.AddBlock("Second block")
    bc.AddBlock("Third block")

    for _, block := range bc.Blocks {
        fmt.Printf("Previous hash: %x\n", block.PrevBlockHash)
        fmt.Printf("Data: %s\n", block.Data)
        fmt.Printf("Created on: %s\n", time.Unix(block.Timestamp, 0))
        fmt.Printf("Hash: %x\n", block.Hash)
        fmt.Println()
    }
}
