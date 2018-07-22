package main

import (
	"log"

	"github.com/tinyzimmer/parity-go"
	"github.com/tinyzimmer/parity-go/eth"
)

func main() {
	node, err := parity.GetParityNode("", true)
	if err != nil {
		log.Fatal(err)
	}
	client := eth.NewClient(node)
	resp, err := client.Syncing()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Encoded: %+v\n", resp)
	log.Printf("Decoded: %+v\n", resp.DecodeAll())
}
