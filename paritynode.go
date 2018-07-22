package parity

const (
	PARITY_DEFAULT_HOST = "http://localhost:8545"
	JSON_RPC_VERSION    = "2.0"
	ID                  = 1
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
	return
}
