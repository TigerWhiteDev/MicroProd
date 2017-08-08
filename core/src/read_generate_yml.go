package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type cgenerate struct {
	Type string `yaml:"type"`
	Lang string `yaml:"lang"`
}

func (c *cgenerate) getConf(file string) *cgenerate {

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
