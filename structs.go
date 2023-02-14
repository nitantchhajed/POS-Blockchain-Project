type POSNetwork struct {
	Blockchain []*Block
	BlockchainHead *Block
	Validators []*Node
}

type Node struct {
	Stake int
	Address string
}

type Block struct {
	Timestamp string
	PrevHash string
	Hash string
	ValidatorAddr string
}