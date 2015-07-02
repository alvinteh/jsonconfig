package jsonconfig

import (
	"flag"
	"io/ioutil"

	"github.com/bitly/go-simplejson"
)

//LoadFromFile returns a JSON structure representation of the configuration from the supplied filename
func LoadFromFile(filename string) *simplejson.Json {
	//Read configuration file
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	//Unmarshal configuration
	config, err := simplejson.NewJson(bytes)

	if err != nil {
		panic(err)
	}

	return config
}

//LoadFromFlag returns a JSON structure representation of the configuration from the supplied filename flag
func LoadFromFlag(flagName string) *simplejson.Json {
	filename := ""
	flag.StringVar(&filename, flagName, "", "")
	flag.Parse()

	return LoadFromFile(filename)
}
