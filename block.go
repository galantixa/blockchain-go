package blockchain_go

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Index is the position of the data record in the blockchain
// Timestamp is automatically determined and is the time the data is written
// BPM or beats per minute, is your pulse rate
// Hash is a SHA256 identifier representing this data record
// PrevHash is the SHA256 identifier of the previous record in the chain

type Block struct {
	Index     int
	Timestamp string
	Hash      string
	PrevHash  string
	BPM       int
}

var Blockchain []Block

func CalculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func GenerateBlock(oldBlock Block, newBlock Block, BPM int) (Block, error) {
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)
	newBlock.BPM = oldBlock.BPM

	return newBlock, nil
}
