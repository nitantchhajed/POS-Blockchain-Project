package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	math "math/rand"
	"time"
)

type POSNetwork struct {
	Blockchain     []*Block
	BlockchainHead *Block
	Validators     []*Node
}

type Node struct {
	Stake   int
	Address string
}

type Block struct {
	Timestamp     string
	PrevHash      string
	Hash          string
	ValidatorAddr string
}

func (n POSNetwork) GenerateNewBlock(Validator *Node) ([]*Block, *Block, error) {
	if err := n.ValidateBlockchain(); err != nil {
		Validator.Stake -= 10
		return n.Blockchain, n.BlockchainHead, err
	}

	currentTime := time.Now().String()

	//After checking if our Blockchain is in tact,
	// we get the current time of the system store it as Timestamp upon
	// instantiation of a new Block. We also attach the Hash of the previous
	// hash which we have easy access to thanks to BlockchainHead.
	// We will then call the method NewBlockHash() on BlockchainHead,
	// and assign the address of the input Node as our validator address.
	newBlock := &Block{
		Timestamp:     currentTime,
		PrevHash:      n.BlockchainHead.Hash,
		Hash:          NewBlockHash(n.BlockchainHead),
		ValidatorAddr: Validator.Address,
	}
	//Once the fields of the new Block have been filled,
	// we call ValidateBlockCandidate() on the new Block and
	//see if there are any errors. If there are,
	//we penalize the culprit Node and return an error.
	// If everything went fine, we append the new block to our Blockchain.
	//The return statement on line 22 is just a default if we didn’t catch an error.
	if err := n.ValidateBlockCandidate(newBlock); err != nil {
		Validator.Stake -= 10
		return n.Blockchain, n.BlockchainHead, err
	} else {
		n.Blockchain = append(n.Blockchain, newBlock)
	}
	return n.Blockchain, newBlock, nil
}

//We have two functions to accomplish the task of creating
// a unique Hash for each Block. The function NewBlockHash()
//simply takes all of the info of the Block
// and concatenates it into a single string to be passed to newHash().

func NewBlockHash(block *Block) string {
	blockInfo := block.Timestamp + block.Hash + block.PrevHash + block.ValidatorAddr
	return newHash(blockInfo)
}

//newHash() will leverage the crypto/sha256 package to create
// a new SHA256 object stored as h. We then convert the input
//string s into a byte array and write that into h. Finally we call h.
//Sum() to get h into a format where we can call hex.
//EncodeToString() so that we have a string as our final output.

func newHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

//Here we attach the ValidateBlockchain() method to our PoSNetwork
//struct and return a possible error. If the blockchain is empty or
// just has a single block,
// we have no way to make sure it’s correct so we just return nil for the error.

func (n POSNetwork) ValidateBlockchain() error {
	if len(n.Blockchain) <= 1 {
		return nil
	}
	currBlockIdx := len(n.Blockchain) - 1
	prevBlockIdx := len(n.Blockchain) - 2

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

//To add a new Node to our PoSNetwork, we call NewNode()
//which takes in the initial stake of the Node and returns a new array of Node references.

func (n POSNetwork) NewNode(stake int) []*Node {
	newNode := &Node{
		// we just append to our n.Validators array
		//and call randAddress() to generate a unique address for the new Node.

		Stake:   stake,
		Address: randAddress(),
	}
	n.Validators = append(n.Validators, newNode)
	return n.Validators
}

func randAddress() string {
	b := make([]byte, 16)
	_, _ = math.Read(b)
	return fmt.Sprintf("%x", b)
}

//In the SelectWinner() method, we first find the total stake that is
//held within the network by ranging over n.Validators.

// We also add any nodes with a stake greater than zero to the
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

//we need to instantiate a new Proof of Stake network with what’s
// known as the Genesis block, a.k.a. the first block in the blockchain.
// Once we do so, we also set the network’s BlockchainHead equal to that first Block.

func main() {
	// set random seed
	math.Seed(time.Now().UnixNano())

	//generate an Initial PoS Network including a blockchain with genesis block.
	genesisTime := time.Now().String()
	pos := &POSNetwork{
		Blockchain: []*Block{
			{
				Timestamp:     genesisTime,
				PrevHash:      "",
				Hash:          newHash(genesisTime),
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
	for i := 0; i < 5; i++ {
		//We pass the newly selected winner to generate a new Block,
		// and print out the total Stake of each Node for each round.
		winner, err := pos.SelectWinner()
		if err != nil {
			log.Fatal(err)
		}
		winner.Stake += 10
		pos.Blockchain, pos.BlockchainHead, err = pos.GenerateNewBlock(winner)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Round ", i)
		fmt.Println("\tAddress:", pos.Validators[0].Address, "-Stake:", pos.Validators[0].Stake)
		fmt.Println("\tAddress:", pos.Validators[1].Address, "-Stake:", pos.Validators[1].Stake)
	}

	//
}
