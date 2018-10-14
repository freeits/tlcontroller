package main

import (
    "log"
    "net"
    "os"
    "time"
    "encoding/binary"
)

type ControllerState struct {
    phase int
    mode int
}

func main(){
    conn, err := net.Dial("udp", "localhost:1053")
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
        state.phase = int(binary.LittleEndian.Uint16(buffer))
        state.mode = int(binary.LittleEndian.Uint16(buffer[2:]))
        log.Printf("Received phase %v. Controller is in %v mode", state.phase, state.mode)
        time.Sleep(pause)
    }
}
