package main

import (
    "github.com/khokhlin/tlcontroller"
    "flag"
)

var configfile string
var addr string

func init(){
    flag.StringVar(&configfile, "config", "./data/controller.xml", "Path to config file")
    flag.StringVar(&addr, "addr", "0.0.0.0:1053", "Address to listen")
    flag.Parse()
}

func main(){
    controller.Run(configfile, addr)
}
