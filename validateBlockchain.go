//Here we attach the ValidateBlockchain() method to our PoSNetwork 
//struct and return a possible error. If the blockchain is empty or
// just has a single block,
// we have no way to make sure it’s correct so we just return nil for the error.


func ( n POSNetwork) ValidateBlockchain() error {
	if len(n.Blockchain) <=1 {
		return nil
	}
	currBlockIdx := len(n.Blockchain)-1
	prevBlockIdx := len(n.Blockchain)-2

for prevBlockIdx >= 0 {
	currBlock := n.Blockchain[currBlockIdx]
	prevBlock := n.Blockchain[prevBlockIdx]

//check is if the Hash of the previous Block is equal to
// what the current Block has stored for it’s previous Hash.

	if currBlock.PrevHash != prevBlock.Hash {
		return errors.New("blockchainhas has inconsistent hashes")
	}

//check to see if at any point a previous Block’s 
//Timestamp is newer than the current Block.

	if currBlock.Timestamp <= prevBlock.Timestamp {
		return errors.New("blockchain has inconsistent timestamps")
	}

//check that if we directly calculate the Hash of the previous Block,
// that we still get back the Hash of the current Block.
// If any of these conditions hold,
// then we return an error since our blockchain is in a tampered state

	if NewBlockHash(prevBlock) != currBlock.Hash {
		return errors.New("blockchain has inconsistent hash generation")
	}

	currBlockIdx--
	prevBlockIdx--
}
return nil
}