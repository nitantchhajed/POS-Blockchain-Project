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

funct newHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}