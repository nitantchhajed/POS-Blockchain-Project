func (n POSNetwork) ValidateBlockCandidate(newBlock *Block) error {
	if n.BlockchainHead.Hash != newBlock.PrevHash {
		return errors.New("blockchain HEAD hash is not equal to new block previous hash")
	}

	if n.BlockchainHead.Timestamp >= newBlock.Timestamp {
		return errors.New("blockchain timestamp is greater than or equal to new block")
	}

	if NewBlockHash(n.BlockchainHead) != newBlock.Hash {
		return errors.New("new block hash of blockchain head does not equal new block hash")
	}
	return nil

}