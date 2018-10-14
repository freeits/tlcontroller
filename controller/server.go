package controller

import (
    "encoding/binary"
    "log"
    "net"
    "os"
    "sync"
)

func serve(pc net.PacketConn, addr net.Addr, buf []byte) {
	pc.WriteTo(buf, addr)
}

type ControllerState struct {
    phase int
    mode int
}

type EventListener struct {
    ch chan int
    state *ControllerState
}

func (l *EventListener) listenEvents() {
    for {
        l.state.phase = <-l.ch
        l.state.mode = <-l.ch
    }
}

func Server(wg *sync.WaitGroup, ch chan int) {
    defer wg.Done()

    state := ControllerState{}
    l := EventListener{ch, &state}
    go l.listenEvents()

    pc, err := net.ListenPacket("udp", ":1053")
	if err != nil {
        log.Print(err)
        os.Exit(1)
	}
	defer pc.Close()

	for {
		buf := make([]byte, 4, 4)
		_, addr, err := pc.ReadFrom(buf)
		if err != nil {
			continue
		}
        binary.LittleEndian.PutUint16(buf, uint16(state.phase))
        binary.LittleEndian.PutUint16(buf[2:], uint16(state.mode))
        log.Printf("Sending data: %v\n", buf)
		go serve(pc, addr, buf)
	}
}
