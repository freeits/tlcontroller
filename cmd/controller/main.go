package main

import (
    "flag"
    "github.com/khokhlin/traffic_lights/controller"
)

var configfile string
var addr string

func init(){
    flag.StringVar(&configfile, "config", "./data/controller.xml", "Path to config file")
    flag.StringVar(&addr, "addr", "0.0.0.0:1053", "Listening address")
    flag.Parse()
}

func main(){
    controller.Run(configfile, addr)
}
