package main

import (
	"github.com/fari-proxy/client/util"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

func main () {
	var conf string
	var config map[string]string
	flag.StringVar(&conf, "c", ".client.json", "client config")
	flag.Parse()

	bytes, err := ioutil.ReadFile(conf)
	if err != nil {
		log.Fatal("read %s failed.", conf)
	}

	if err := json.Unmarshal(bytes, &config); err != nil {
		log.Fatal("parse %s failed.", conf)
	}
	client := client.NewClient(config["remote_addr"], config["listen_addr"], config["password"])
	client.Listen()
}
