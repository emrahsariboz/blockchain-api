package blockchain

//block defines data format for blockchain
//swagger:model
type block struct {
	//Hash of previous block
	//
	//required: true
	PrevHash    string
	CurrentHash string
	Nonce       int32
	Timestamp   string
}

type blockchain struct {
	Blocks []block
}

func (b *blockchain) GetBlocks() []block {
	return b.Blocks
}

func NewBlockchain() *blockchain {
	return &blockchain{}
}
