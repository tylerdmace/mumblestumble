package network

// Operations -
//
// 0x00: Version
// 0x01: Version Ack
// 0x02: Ping
// 0x03: Pong
// 0x04: Get Blocks
// 0x05: Send Blocks
// 0x06: Get Txns
// 0x07: Send Txns
// 0x08: Get Addrs
// 0x09: Send Addrs

func SendOp(node Node, op byte) {
	_ = GetLocalAddresses()
}
