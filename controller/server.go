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
    phase uint16
    mode uint16
    mutex *sync.Mutex
}

func NewControllerState() *ControllerState{
    mutex := &sync.Mutex{}
    return &ControllerState{mutex: mutex}
}

type Server struct {
    addr string
}

func makeServer(addr string) Server {
    return Server{addr}
}

func (s *Server) serve(wg *sync.WaitGroup, state *ControllerState) {
    defer wg.Done()

    log.Printf("Server is listening on %v", s.addr)

    pc, err := net.ListenPacket("udp", s.addr)
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
        state.mutex.Lock()
        binary.LittleEndian.PutUint16(buf, state.phase)
        binary.LittleEndian.PutUint16(buf[2:], state.mode)
        state.mutex.Unlock()
        log.Printf("Sending data: %v\n", buf)
		go serve(pc, addr, buf)
	}
}
