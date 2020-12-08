package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Node struct {
	Name    string
	Address string
}

func NewNode(name string, address string) error {
	return nil
}

// func (n *Node)

func (n *Node) sendRequest(requestBody []byte) ([]byte, error) {
	resp, err := http.Post(n.Address, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
