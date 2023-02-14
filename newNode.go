//To add a new Node to our PoSNetwork, we call NewNode()
//which takes in the initial stake of the Node and returns a new array of Node references.

func (n POSNetwork) NewNode(stake int) []*Node {
	newNode := &Node {
// we just append to our n.Validators array 
//and call randAddress() to generate a unique address for the new Node.

		Stake: stake,
		Address: randAddress(),
	}
	n.Validators = append(n.Validators, newNode)
	return n.Validators
}

func randAddress() string {
	b:= make([]byte, 16)
	_, _ = math.Read(b)
	return fmt.Sprintf("%x", b)
}