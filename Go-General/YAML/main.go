package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ertugrulbal/goconfigop/models"
	"gopkg.in/yaml.v2"
)

/*
	YAML
	- http://gopkg.in/yaml.v2
	- http://godoc.org/gopkg.in/yaml.v2
*/

// Konfigürasyon Dosyaları

// Araştırma Konusu : Secret Management Tools

func main() {

	// v1

	fileName := "config.yaml"
	source, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(source))

	// v2

	var config models.Config
	yaml.Unmarshal(source, &config)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(config)
	// fmt.Println(config.Settings)
	// fmt.Println(config.Port)
	// fmt.Println(config.Settings[0])

	for _, v := range config.Settings {
		fmt.Println("Ayar : " + v)
	}
}
