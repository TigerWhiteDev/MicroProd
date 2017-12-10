package readyml

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Lang struct {
	Vendor string `yaml:"vendor"`
}

func (c *Lang) GetConf(file string) *Lang {

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
