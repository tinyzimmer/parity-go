package parityrpc

import (
	"log"
	"net"
	"strings"
	"time"
)

const (
	DEFAULT_TEST_TIMEOUT = 2
	PARITY_DEFAULT_HOST  = "http://localhost:8545"
	JSON_RPC_VERSION     = "2.0"
	ID                   = 1
)

type ParityNode struct {
	Host      string
	Debug     bool
	IdCounter uint64
}

func GetParityNode(url string, debug bool) (n ParityNode, err error) {

	// Configure a struct representation of a Parity Node and test connectivity

	if url == "" {
		url = PARITY_DEFAULT_HOST
	}
	if debug {
		log.Printf("DEBUG: Initializing Parity client for %s\n", url)
	}
	n = ParityNode{url, debug, 1}
	err = n.TestConnection()
	return

}

func (n *ParityNode) TestConnection() (err error) {

	// Dies a simple socket bind test to the jsonrpc endpoint of the Parity node
	// This confirms whether or not we have access to it.

	if n.Debug {
		log.Printf("DEBUG: Testing connectivity to %s\n", n.Host)
		log.Printf("DEBUG: Using built-in timeout of %v seconds\n", DEFAULT_TEST_TIMEOUT)
	}
	stripHttp := strings.Replace(n.Host, "http://", "", -1)
	con, err := net.DialTimeout("tcp", stripHttp, time.Duration(DEFAULT_TEST_TIMEOUT*time.Second))
	if err != nil {
		log.Printf("ERROR: Could not reach %s\n", n.Host)
		return
	} else {
		if n.Debug {
			log.Printf("DEBUG: %s is available. Closing test connection.\n", n.Host)
		}
		con.Close()
	}
	return

}
