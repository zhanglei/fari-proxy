package main

import (
	"github.com/fari-proxy/server/util"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

func main () {
	var conf string
	var config map[string]string
	flag.StringVar(&conf, "c", ".server.json", "server config")
	flag.Parse()

	bytes, err := ioutil.ReadFile(conf)
	if err != nil {
		log.Fatal("read %s failed.", conf)
	}

	if err := json.Unmarshal(bytes, &config); err != nil {
		log.Fatal("parse %s failed.", conf)
	}
	server := server.NewServer(config["listen_addr"], config["password"])
	server.Listen()
}
