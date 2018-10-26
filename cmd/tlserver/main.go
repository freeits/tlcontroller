package main

import (
    "flag"
    "github.com/khokhlin/traffic_lights/tlserver"
)

var host string
var port int

func init(){
    flag.StringVar(&host, "host", "0.0.0.0", "Controller host")
    flag.IntVar(&port, "port", 1053, "Controller port")
    flag.Parse()
}

func main(){
    tlserver.Run(host, port)
}
