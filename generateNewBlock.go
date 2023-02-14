

func (n POSNetwork) GenerateNewBlock(Validator *Node) ([]*Block, *Block, error) {
	if err := n.ValidateBlockchain(); err != nil {
		Validator.Stake -=10
		return n.Blockchain, n.BlockchainHead, err
	}

	currentTime := time.Now().String()

//After checking if our Blockchain is in tact,
// we get the current time of the system store it as Timestamp upon
// instantiation of a new Block. We also attach the Hash of the previous
// hash which we have easy access to thanks to BlockchainHead.
// We will then call the method NewBlockHash() on BlockchainHead,
// and assign the address of the input Node as our validator address.
	newBlock := &Block {
		Timestamp: currentTime,
		PrevHash: n.BlockchainHead.Hash,
		Hash: NewBlockHash(n.BlockchainHead),
		ValidatorAddr: Validator.Address,
	}
//Once the fields of the new Block have been filled,
// we call ValidateBlockCandidate() on the new Block and 
//see if there are any errors. If there are, 
//we penalize the culprit Node and return an error.
// If everything went fine, we append the new block to our Blockchain. 
//The return statement on line 22 is just a default if we didn’t catch an error.
	if err := n.ValidateBlockCandidate(newBlock); err != nil {
		Validator.Stake -=10
		return n.Blockchain, n.BlockchainHead, err
	} else {
			n.Blockchain = append(n.Blockchain, newBlock)
	}
	return n.Blockchain, newBlock, nil
}