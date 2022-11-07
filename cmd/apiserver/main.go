package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/artemxgod/http-rest-api/internal/app/apiserver"
)

var configPath string

/* setting flag for our execution file.
	Flag usage example: ./appserver -config-path ./configs/second.toml
 	will change config file from default to second.toml */
func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to cfg file")
}

func main() {
	// parsing all flags
	flag.Parse()

	// defining default config
	config := apiserver.NewConfig()
	// getting config from file
	_, err := toml.DecodeFile(configPath, config)
	FatalOnErr(err)

	// starting our http server
	if err := apiserver.Start(config); err != nil {
		log.Fatal("ERROR -", err)
	}
}

func FatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}