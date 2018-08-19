package main

import (
	"log"
	"time"

	"github.com/tinyzimmer/parity-go"
	"github.com/tinyzimmer/parity-go/parity"
)

// Example command for starting Parity
//docker run -ti -p 8180:8180 -p 8545:8545 -p 8546:8546 -p 30303:30303 -p 30303:30303/udp \
//    --rm parity/parity --no-ancient-blocks --no-serve-light --max-peers 250 \
//    --snapshot-peers 50 --min-peers 50 --mode active --tracing off --pruning fast \
//    --db-compaction ssd --cache-size 4096 --chain kovan --jsonrpc-interface all

func main() {

	// Connect to a Parity Node. We take the default localhost:8545 and
	// use debug to see all output

	node, err := parityrpc.GetParityNode("", false)
	if err != nil {
		log.Fatal(err)
	}

	// Create a client for the eth api

	client := parity.NewClient(node)

	for {

		// query the sync status of the node

		resp, err := client.NetPeers()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Response: %+v\n", resp)

		time.Sleep(time.Duration(3 * time.Second))

	}
}
