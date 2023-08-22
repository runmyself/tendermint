package l2

import "github.com/tendermint/tendermint/types"

var _ L2node = &MockL2Node{}

type MockL2Node struct {
	height int64
}

func NewMockL2Node(h int64) L2node {
	return &MockL2Node{
		height: h,
	}
}

func (ml *MockL2Node) FetchBlock(height int64) (txs [][]byte, l2Config, zkConfig []byte, err error) {
	txs = [][]byte{[]byte("a"), []byte("b"), []byte("c")}
	l2Config = []byte("d")
	zkConfig = []byte("e")
	return
}

func (ml *MockL2Node) DeliverBlock(
	txs [][]byte,
	l2Config []byte,
	zkConfig []byte,
	validators []types.Address,
	blsSignatures [][]byte,
) (
	err error,
) {
	return
}
