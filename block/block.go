package block

import (
	"github.com/tylerdmace/mumblestumble/txn"
)

type Block struct {
	Header BlockHeader
}

type BlockHeader struct {
	Height        []byte
	Version       []byte
	PreviousBlock []byte
	Timestamp     []byte
	Bits          []byte
	Nonce         []byte
	Transactions  []txn.Txn
}

func GenerateGenesis() Block {
	var gb Block

	gb.Header.Height = []byte{}
	gb.Header.Version = []byte{}
	gb.Header.PreviousBlock = []byte{}

	return gb
}
