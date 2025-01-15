package network

import (
	"blockchain-go/blockchain"
	"encoding/json"
	"log"
	"net"
)

func Broadcast(addresses []string, data []blockchain.Block) {
	for _, address := range addresses {
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Println("Error connecting to node:", err)
			continue
		}
		defer conn.Close()
		encoder := json.NewEncoder(conn)
		err = encoder.Encode(data)
		if err != nil {
			log.Println("Error encoding data:", err)
		} else {
			log.Println("Broadcast success")
		}
	}
}
