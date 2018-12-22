package tlserver

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type ControllerState struct {
	phase uint16
	mode  uint16
}

func Run(host string, port int) {
	addr := fmt.Sprintf("%v:%v", host, port)
	conn, err := net.Dial("udp", addr)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	defer conn.Close()

	buffer := make([]byte, 4)
	pause := time.Duration(1) * time.Second
	for {
		conn.Write([]byte("1"))
		conn.Read(buffer)
		state := ControllerState{}
		state.phase = binary.LittleEndian.Uint16(buffer)
		state.mode = binary.LittleEndian.Uint16(buffer[2:])
		log.Printf("Received phase %v. Controller is in %v mode", state.phase, state.mode)
		time.Sleep(pause)
	}
}
