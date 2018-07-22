package parity

import (
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
	Host  string
	Debug bool
}

func GetParityNode(url string, debug bool) (n ParityNode, err error) {
	if url == "" {
		url = PARITY_DEFAULT_HOST
	}
	n = ParityNode{url, debug}
	err = n.TestConnection()
	return
}

func (n *ParityNode) TestConnection() (err error) {
	stripHttp := strings.Replace(n.Host, "http://", "", -1)
	con, err := net.DialTimeout("tcp", stripHttp, time.Duration(DEFAULT_TEST_TIMEOUT*time.Second))
	if err != nil {
		return
	} else {
		con.Close()
	}
	return
}
