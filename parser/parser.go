package parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// ConfigStruct contains the struct representation of the config file.
type ConfigStruct struct {
	Controller struct{}
	Logger     struct{}
	Main       struct {
		TLSCert string
		TLSKey  string
	}
	Model struct{}
	View  struct{}
}

// Config provides a struct with settings for other modules. Each part of
// the configStruct should be given to the corresponding module.
func Config(path string) *ConfigStruct {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var config ConfigStruct

	if err := json.Unmarshal(buf, &config); err != nil {
		log.Fatal(err)
	}

	return &config
}
