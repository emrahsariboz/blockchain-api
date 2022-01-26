package blockchain

type block struct {
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
