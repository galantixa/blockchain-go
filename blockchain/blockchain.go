package blockchain

// Blockchain is a slice of Block
var Blockchain []Block

// replaceChain replaces the current blockchain with a new one if it's longer
func replaceChain(newBlocks []Block) []Block {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
	return Blockchain
}

// Node structure represents a node in the blockchain network
type Node struct {
	Address    string
	Blockchain []Block
}

func addPeer(node Node, address string) {
	// updated later
}
