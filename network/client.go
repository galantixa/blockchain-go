package network

import (
	"blockchain-go/blockchain"
	"encoding/json"
	"log"
	"net"
)

func SendData(address string, data []blockchain.Block) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Println("Error connecting to node:", err)
		return
	}
	defer conn.Close()

	encoder := json.NewEncoder(conn)
	err = encoder.Encode(data)
	if err != nil {
		log.Println("Error encoding data:", err)
	} else {
		log.Println("Data sent to", address)
	}
}
