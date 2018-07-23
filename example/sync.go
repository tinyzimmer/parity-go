package main

import (
	"log"
	"time"

	"github.com/tinyzimmer/parity-go"
	"github.com/tinyzimmer/parity-go/eth"
)

func main() {

	// Connect to a Parity Node. We take the default localhost:8545 and
	// use debug to see all output

	node, err := parity.GetParityNode("", true)
	if err != nil {
		log.Fatal(err)
	}

	// Create a client for the eth api

	client := eth.NewClient(node)

	for {

		// query the sync status of the node

		resp, err := client.Syncing()
		if err != nil {
			log.Fatal(err)
		}

		// Log the hexadecimal encoded responses and decode them as well.

		log.Printf("Encoded: %+v\n", resp)
		log.Printf("Decoded: %+v\n", resp.DecodeAll())

		time.Sleep(time.Duration(3 * time.Second))

	}
}
