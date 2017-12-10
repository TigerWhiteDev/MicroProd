package readyml

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Project struct {
	Type string `yaml:"type"`
	Lang string `yaml:"lang"`
	Git  string `yaml:"git"`
}

func (c *Project) GetConf(file string) *Project {

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
