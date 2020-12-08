package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	nodeList map[string]*Node
)

type Node struct {
	Name            string
	RPCPort         int
	DBPath          string
	LogPath         string
	IsValidatorNode bool
	IsPNode         bool
}

type NodeInfo struct {
	MachineInfo struct {
		Mem int
		CPU int
	}
	ChainInfo struct {
		BeaconHeight int
		ShardsHeight map[int]int
		DiskUsed     int
	}
	IsValidatorNode bool
	IsPNode         bool
}

func NewNode(name string, address string) error {
	return nil
}

func (n *Node) GetInfo() {}

func (n *Node) sendRequest(requestBody []byte) ([]byte, error) {
	resp, err := http.Post(fmt.Sprintf("127.0.0.1:+%v", n.RPCPort), "application/json", bytes.NewBuffer(requestBody))
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
