package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	nodeList map[string]*Node
)

type Node struct {
	Name            string
	Addresss        string
	DBPath          string
	LogPath         string
	IsValidatorNode bool `yaml:"ValidatorNode"`
	IsPNode         bool `yaml:"PNode"`
	status          NodeStatus
}

type NodeStatus struct {
	MachineInfo struct {
		MemUsed uint64
		CPU     float64
	}
	ChainInfo struct {
		BeaconHeight uint64
		ShardsHeight map[byte]uint64
		DiskUsed     float64
	}
	IsValidatorNode bool
	IsPNode         bool
}

type ErrMsg struct {
	Code       int
	Message    string
	StackTrace string
}

func (n *Node) GetStatus() NodeStatus {
	return n.status
}

func (n *Node) StartMonitor() {
	for {
		time.Sleep(3 * time.Second)
		beacon, shards, err := n.getInfo()
		if err != nil {
			log.Println(err)
			continue
		}
		disks := dirSizeMB(n.DBPath)
		// fmt.Printf("disks value: %v\n", disks)
		n.status.ChainInfo.DiskUsed = disks
		n.status.MachineInfo.CPU = system.CPU
		n.status.MachineInfo.MemUsed = system.MemUsed
		n.status.ChainInfo.BeaconHeight = beacon
		n.status.ChainInfo.ShardsHeight = shards
	}
}

func (n *Node) getInfo() (uint64, map[byte]uint64, error) {
	requestBody, rpcERR := json.Marshal(map[string]interface{}{
		"jsonrpc": "1.0",
		"method":  "getblockchaininfo",
		"params":  []interface{}{},
		"id":      1,
	})
	if rpcERR != nil {
		return 0, nil, rpcERR
	}
	body, err := n.sendRequest(requestBody)
	if err != nil {
		return 0, nil, err
	}
	resp := struct {
		Result GetBlockChainInfoResult
		Error  *ErrMsg
	}{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return 0, nil, errors.New(rpcERR.Error())
	}
	if resp.Error != nil && resp.Error.StackTrace != "" {
		return 0, nil, errors.New(resp.Error.StackTrace)
	}
	beaconHeight := resp.Result.BestBlocks[-1].Height
	shardsHeight := make(map[byte]uint64)

	for shardID, block := range resp.Result.BestBlocks {
		if shardID != -1 {
			shardsHeight[byte(shardID)] = block.Height
		}
	}
	return beaconHeight, shardsHeight, err
}

func (n *Node) sendRequest(requestBody []byte) ([]byte, error) {
	resp, err := http.Post(n.Addresss, "application/json", bytes.NewBuffer(requestBody))
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

func dirSizeMB(path string) float64 {
	var dirSize int64 = 0
	readSize := func(path string, file os.FileInfo, err error) error {
		if !file.IsDir() {
			dirSize += file.Size()
		}

		return nil
	}
	filepath.Walk(path, readSize)
	sizeMB := float64(dirSize) / 1024.0 / 1024.0
	return sizeMB
}
