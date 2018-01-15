package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// type for the JSON file values
type Configuration struct {
	Port          string
	FolderToServe string
}

func main() {
	// read the command line params
	configFileLocation := flag.String("c", "/", "config file location to use")
	configFileName := flag.String("f", "config.json", "Filename for config, must be json format")
	flag.Parse()
	// create the config path variable
	configFile := *configFileLocation + "/" + *configFileName

	// return the config to variables
	port, directory := readConfig(configFile)

	fs := http.FileServer(http.Dir(directory))
	http.Handle("/", fs)

	log.Printf("Serving %s on HTTP port: %s\n", directory, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func readConfig(fileName string) (string, string) {
	// outputs the fileName to the console
	log.Printf(fileName)
	// opens the config file at the location passed in
	file, _ := os.Open(fileName)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	// decodes the JSON file into the type
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	// return the config file values
	return configuration.Port, configuration.FolderToServe
}

// /Users/johnspurgin/Documents/Git/Grid/Arthaus
