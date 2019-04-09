package chain

import "fmt"

type Chain struct {
	ID []byte
}

func NewChain() Chain {
	var c Chain
	return c
}

func (c *Chain) PrintInfo() {
	fmt.Printf("Chain ID: %v\r\n", c.ID)
}
