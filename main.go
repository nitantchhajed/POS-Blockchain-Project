//we need to instantiate a new Proof of Stake network with what’s
// known as the Genesis block, a.k.a. the first block in the blockchain.
// Once we do so, we also set the network’s BlockchainHead equal to that first Block.





func main() {
	// set random seed
	math.Seed(time.Now().UnixNano())

	//generate an Initial PoS Network including a blockchain with genesis block.
	genesisTime := time.Now().String()
	pos := &PoSNetwork {
		Blockchain: []*Block {
			{
				Timestamp: genesisTime,
				PrevHash: "",
				Hash: newHash(genesisTime),
				ValidatorAddr: "",
			},
		},
	}

	pos.BlockchainHead = pos.Blockchain[0]

//we add two Nodes to the network to act as validators
// with 60 and 40 tokens as their initial Stake. 

//instantiate nodes to act as validators in our network
	pos.Validators = pos.NewNode(60)
	pos.Validators = pos.NewNode(40)

//build 5 additions to the blockchain
//For five iterations we will select a new winner for our
// Blockchain and crash our program if there’s any errors — because prototyping.
for i:=0; i < 5; i++ {
	//We pass the newly selected winner to generate a new Block,
	// and print out the total Stake of each Node for each round.
	winner, err := pos.SelectWinner()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Round", i)
	fmt.Println("\tAddress:", pos.Validators[0].Address, "-Stake:", pos.Validators[0].Stake)
	fmt.Println("\tAddress:", pos.Validators[1].Address,"-Stake", pos.validators[1].Stake)
}
	pos.PrintBlockcahinInfo()
}