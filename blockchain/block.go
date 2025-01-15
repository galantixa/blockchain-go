package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Transaction data

type Transaction struct {
	From      string  `json:"from"`
	To        string  `json:"to"`
	Amount    float64 `json:"amount"`
	Timestamp string  `json:"timestamp"`
}

// Block structure
type Block struct {
	Index        int           `json:"index"`
	Timestamp    string        `json:"timestamp"`
	Hash         string        `json:"hash"`
	PrevHash     string        `json:"prevHash"`
	Transactions []Transaction `json:"transactions"`
	BPM          int           `json:"bpm"`
}

// Calculate hash from block
func CalculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// Create Block
func GenerateBlock(oldBlock Block, newBlock Block, BPM int, transactions []Transaction) (Block, error) {
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)
	newBlock.Transactions = transactions
	newBlock.BPM = BPM // Set the BPM for the new block

	return newBlock, nil
}

// Initialize Genesis block
func GenerateGenesisBlock() Block {
	transactions := []Transaction{
		{
			From:      "Prometheus",
			To:        "address1",
			Amount:    100.0,
			Timestamp: time.Now().String(),
		},
	}
	return Block{
		Index:        0,
		Timestamp:    time.Now().String(),
		Transactions: transactions,
		PrevHash:     "",
		Hash: CalculateHash(Block{
			Index:        0,
			Timestamp:    time.Now().String(),
			Transactions: transactions,
			PrevHash:     "",
		}),
	}

}

// ValidateBlock checks if a new block is valid
func ValidateBlock(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}
