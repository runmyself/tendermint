package l2

import "github.com/tendermint/tendermint/types"

type L2node interface {
	FetchBlock(height int64) (txs [][]byte, l2Config, zkConfig []byte, err error)
	DeliverBlock(txs [][]byte, l2Config, zkConfig []byte, validators []types.Address, blsSignatures [][]byte) (err error)
}

func GetValidators(commit *types.Commit) (validators []types.Address) {
	for _, v := range commit.Signatures {
		validators = append(validators, v.ValidatorAddress)
	}
	return
}

func GetBLSSignatures(commit *types.Commit) (blsSignatures [][]byte) {
	for _, v := range commit.Signatures {
		blsSignatures = append(blsSignatures, v.BLSSignature)
	}
	return
}

func ConvertBytesSliceToTxs(txs [][]byte) []types.Tx {
	s := make([]types.Tx, len(txs))
	for i, v := range txs {
		s[i] = v
	}
	return s
}
