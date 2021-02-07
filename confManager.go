package main

import (
	"encoding/json"
	"os"
	"fmt"
)

type Configuration struct{
	WebSiteAnalyse string
	LogFile string
	ForbiddenExtensions []string
}

func (conf *Configuration) readConf() Configuration{
	file, _ := os.Open("configuration.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("Impossible to get the configuration")
	}
	return configuration
}