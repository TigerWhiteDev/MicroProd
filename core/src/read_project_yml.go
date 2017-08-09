package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type cproject struct {
	Type string `yaml:"type"`
	Lang string `yaml:"lang"`
}

func (c *cproject) getConf(file string) *cproject {

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
