package parity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type PostPayload struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      int         `json:"id"`
	JsonRPC string      `json:"jsonrpc"`
}

type ParityResponse struct {
	JsonRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
}

func (c *ParityNode) Post(method string, params interface{}) (response ParityResponse, success bool, err error) {
	payload, err := c.GenPayload(method, params)
	if err != nil {
		return
	}
	req, err := GenRequest(c.Host, payload)
	if err != nil {
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	json.Unmarshal(data, &response)
	if string(response.Result) == "false" {
		success = false
	} else {
		success = true
	}
	return
}

func (c *ParityNode) GenPayload(method string, params interface{}) (payload []byte, err error) {
	pstruct := PostPayload{
		Method:  method,
		Params:  params,
		Id:      1,
		JsonRPC: JSON_RPC_VERSION,
	}
	payload, err = json.Marshal(pstruct)
	return
}

func GenRequest(host string, payload []byte) (req *http.Request, err error) {
	req, err = http.NewRequest("POST", host, bytes.NewBuffer(payload))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	return
}

func HexToInt(hex interface{}) uint64 {
	hexStr := fmt.Sprintf("%v", hex)
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	return uint64(result)
}
