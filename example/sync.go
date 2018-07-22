package main

import (
	"log"

	parity "github.com/tinyzimmer/parity-go"
	"github.com/tinyzimmer/parity-go/eth"
)

func main() {
	node, err := parity.GetParityNode("", false)
	if err != nil {
		log.Fatal(err)
	}
	client := eth.NewClient(node)
	resp, err := client.BlockNumber()
	if err != nil {
		return
	}
	log.Printf("%+v\v", resp)
}
