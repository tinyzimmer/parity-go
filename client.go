package parity

import (
	"log"
)

const (
	PARITY_DEFAULT_HOST = "http://localhost:8545"
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

func main() {
	_, err := GetParityNode("", false)
	if err != nil {
		log.Fatal(err)
	}
}
