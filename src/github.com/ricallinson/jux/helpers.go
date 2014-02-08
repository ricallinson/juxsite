package jux

import (
	"github.com/ricallinson/forgery"
	"io/ioutil"
	"launchpad.net/goyaml"
)

// Creates a new site instance and returns it.
func New() *Site {
	site := &Site{}
	return site
}

// Returns the jux.Site from the given req.Map["juxsite"].
func GetSite(req *f.Request) *Site {
	return req.Map[JuxSite].(*Site)
}

// Returns the Layout Map position for the current component.
func GetPosition(req *f.Request) string {
	return req.Params["fcomposite_id"][0:PositionSize]
}

// Return the given interface as a YAML byte slice.
func ToYaml(i interface{}) []byte {
	data, err1 := goyaml.Marshal(i)
	if err1 != nil {
		panic(err1)
		return []byte{}
	}
	return data
}

// Reads the given YAML byte slice into the given interface.
func FromYaml(yaml []byte, i interface{}) {
	// Unmarshal the source into this Config instance.
	err2 := goyaml.Unmarshal(yaml, i)
	if err2 != nil {
		panic(err2)
		return
	}
}

// Reads the given file and return a byte slice.
func FromFile(filepath string) []byte {
	// Read the source file.
	file, err1 := ioutil.ReadFile(filepath)
	if err1 != nil {
		// panic(err1)
		return []byte("Not found.")
	}
	return file
}

// Reads the given YAML filepath into the given interface.
func FromYamlFile(filepath string, i interface{}) {
	FromYaml(FromFile(filepath), i)
}
