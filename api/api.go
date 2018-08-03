package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var network = "mainnet"

// Receive takes a handle and returns the
// Response struct with the address.
func Receive(handle string) (*Response, error) {
	handle = strings.TrimLeft(handle, "$")

	url := fmt.Sprintf("%s/receivingAddress/%s", GetURL(), handle)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := &Response{}

	if len(body) == 0 {
		r.Error = "not found"
		return r, nil
	}

	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// GetURL gets the appropriate URL for the network.
func GetURL() string {
	switch network {
	case "mainnet":
		return "https://api.handcash.io/api"
	case "testnet":
		return "https://test-api.handcash.io/api"
	}

	return ""
}

// SetNetwork allows you to set the network.
// Valid values are mainnet/testnet.
func SetNetwork(n string) error {
	switch n {
	case "mainnet":
		network = "mainnet"
	case "testnet":
		network = "testnet"
	default:
		return fmt.Errorf("invalid network - %s", n)
	}

	return nil
}
