package blockchain_go

func ValidateBlock(newBlock, OldBlock Block) bool {
	if OldBlock.Index+1 != newBlock.Index {
		return false
	}

	if OldBlock.Hash != newBlock.Hash {
		return false
	}
	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func replaceChain(newBlocks []Block) []Block {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
	return Blockchain
}
