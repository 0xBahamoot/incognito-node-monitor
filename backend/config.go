package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type ConfigYaml struct {
	APIPort int
	Pass    string
	Nodes   []Node
}

func readYaml() {
	cfg := ConfigYaml{}
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(yamlFile), &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", cfg)
}
