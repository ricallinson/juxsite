package jux

import (
	"github.com/ricallinson/forgery"
	"io/ioutil"
	"launchpad.net/goyaml"
)

// Returns the jux.Config from the given req.Map["cfg"].
func GetConfig(req *f.Request) *Config {
	return req.Map["cfg"].(*Config)
}

// Return the given interface as a YAML string.
func ToYaml(i interface{}) []byte {
	data, err1 := goyaml.Marshal(i)
	if err1 != nil {
		panic(err1)
	}
	return data
}

// Read the given YAML into the given interface.
func FromYaml(yaml []byte, i interface{}) {
	// Unmarshal the source into this Config instance.
	err2 := goyaml.Unmarshal(yaml, i)
	if err2 != nil {
		panic(err2)
		return // error
	}
}

// Read the given YAML file into the given interface.
func FromYamlFile(filepath string, i interface{}) {
	// Read the source file.
	yaml, err1 := ioutil.ReadFile(filepath)
	if err1 != nil {
		panic(err1)
		return // error
	}
	FromYaml(yaml, i)
}
