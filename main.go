package main

import (
	"blockchain-go/blockchain"
	"blockchain-go/network"
	"os"
)

func main() {
	genesisBlock := blockchain.GenerateGenesisBlock()
	node := blockchain.Node{Address: ":8080", Blockchain: []blockchain.Block{genesisBlock}}

	go network.StartServer(node)

	if len(os.Args) > 1 {
		network.SendData(os.Args[1], node.Blockchain)
	}
	select {}
}
