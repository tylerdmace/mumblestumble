package network

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/tylerdmace/mumblestumble/config"
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

func Bootstrap(cfg config.Config) {
	// First, bootstrap the network by starting a listener and connecting to seed nodes
	ln, err := net.Listen("tcp", fmt.Sprintf("%v:%v", cfg.LocalAddress, cfg.LocalPort))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listening at %v:%v\r\n", cfg.LocalAddress, cfg.LocalPort)

	seeds := GetSeeds((byte)(cfg.Network)) // TODO: Support DNS seed servers

	for s := range seeds {
		// Check connectivity of seed nodes
		SendOp(seeds[s], 0x02)
	}

	// TODO: Then download chain state from all pooled peers

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Printf("New connection! (%v)\r\n", conn.RemoteAddr())
}

func GetSeeds(network byte) []Node {
	// TODO: Support DNS seeding
	var seeds []Node
	var s1, s2 Node

	if network == 0x00 { // Mainnet seeds
		s1 = NewNode("192.168.1.201", 25519, 0x00)
		s2 = NewNode("192.168.1.202", 25519, 0x00)
	} else { // Single testnet catch-all for now
		s1 = NewNode("192.168.1.201", 25520, 0x01)
		s2 = NewNode("192.168.1.202", 25520, 0x01)
	}

	seeds = append(seeds, s1)
	seeds = append(seeds, s2)

	return seeds
}

func GetLocalAddresses() []net.IP {
	// TODO: Update to determine local addresses automagically
	var addrs []net.IP

	return addrs
}
