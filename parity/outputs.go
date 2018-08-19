package parity

type ModeOutput string

type EnodeOutput string

type NetPeersOutput struct {
	Active    uint64              `json:"active"`
	Connected uint64              `json:"connected"`
	Max       uint64              `json:"max"`
	Peers     []map[string]string `json:"peers"`
}
