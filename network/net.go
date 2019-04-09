package network

import (
	"net"
	"time"
)

// Our network message protocol -
//
// Network: 1 byte
// Opcode: 1 byte
// Payload Length: 4 bytes
// Payload Checksum: 4 bytes
// Payload: n bytes (variable)

type Peer struct {
	node     Node
	lastSeen time.Time
}

type Node struct {
	addr    net.IP
	port    int  // 25519: Mainnet, 25520: Testnet (can be configured differently per node)
	network byte // 0x00: Mainnet, 0x01: Testnet
}

func NewNode(ip string, port int, network byte) Node {
	n := new(Node)
	n.addr = net.ParseIP(ip)
	n.port = port
	n.network = network

	return *n
}

func GetSeeds(network byte) []Node {
	// Note: Eventually we wanna bootstrap via hardcoded DNS servers
	var seeds []Node
	var s1, s2 Node

	if network == 0x00 { // Mainnet seeds
		s1 = NewNode("192.168.1.201", 25519, 0x00)
		s2 = NewNode("192.168.1.202", 25519, 0x00)
	} else { // Testnets will need to be broken out here; for now a catch-all to a single testnet
		s1 = NewNode("192.168.1.201", 25520, 0x01)
		s2 = NewNode("192.168.1.202", 25520, 0x01)
	}

	// Add nodes to our slice
	seeds = append(seeds, s1)
	seeds = append(seeds, s2)

	return seeds
}

func GetLocalAddresses() []net.IP {
	// TODO: Update to determine local addresses automagically
	var addrs []net.IP

	return addrs
}
