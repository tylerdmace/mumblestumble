package block

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
}

func GenerateGenesis() Block {
	var gb Block

	gb.Header.Height = []byte{}
	gb.Header.Version = []byte{}
	gb.Header.PreviousBlock = []byte{}

	return gb
}
