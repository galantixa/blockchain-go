package network

import (
	"blockchain-go/blockchain"

	"encoding/json"
	"fmt"
	"log"
	"net"
)

func StartServer(node blockchain.Node) {
	listener, err := net.Listen("tcp", node.Address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	fmt.Println("Listening on " + node.Address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go HandleConnetion(conn, node)
	}
}

func HandleConnetion(conn net.Conn, node blockchain.Node) {
	defer conn.Close()
	var receiveBlockchain []blockchain.Block
	decoder := json.NewDecoder(conn)
	err := decoder.Decode(&receiveBlockchain)
	if err != nil {
		log.Println("Error decoding data:", err)
		return
	}
	fmt.Println("Received block chain: ", receiveBlockchain)

}
