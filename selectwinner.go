//In the SelectWinner() method, we first find the total stake that is 
//held within the network by ranging over n.Validators.

//We also add any nodes with a stake greater than zero to the
// array winnerPool for possible selection.
func (n POSNetwork) SelectWinner() (*Node, error) {
	var winnerPool []*Node
	totalStake := 0
	for _, node := range n.Validators {
		if node.Stake > 0 {
			winnerPool = append(winnerPool, node)
			totalStake += node.Stake 
		}
	}
	//If we find winnerPool to be empty, we return an error.
	if winnerPool == nil {
		return nil, errors.New("there are no nodes with stake in the network")
	}
	//Then we select a winning number using the Intn()
	// method which will pick a random number between 0 and our total stake.
	winnerNumber := math.Intn(totalStake)
	tmp := 0

// In order to keep each node having a chance of winning that is proportional
//to it’s total Stake in the network, we incrementally add the Stake of the
// current Node to the tmp variable. If at any point
// the winning number is less than tmp, that Node is selected as our winner.

	for _, node := range n.Validators {
		tmp += node.Stake
		if winnerNumber < tmp {
			return node, nil
		}
	}
	return nil, errors.New("a winner should have been picked but wasn't")
}