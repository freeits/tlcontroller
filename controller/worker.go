package controller

import (
    "log"
    "sync"
    "time"
)

const LOCAL_MODE = 1

type Worker struct {
    cycle Cycle
}

func (m *Worker) work(wg *sync.WaitGroup, state *ControllerState) {
    defer wg.Done()
    log.Print("Worker started")
    for true {
        phase := m.cycle.nextPhase()
        phase_time := time.Duration(phase.time.total_seconds) * time.Second
        time.Sleep(phase_time)
        log.Printf("Phase %v is active", phase.name)
        state.mutex.Lock()
        state.phase = phase.id
        state.mode = LOCAL_MODE
        state.mutex.Unlock()
    }
}

func makeWorker(configfile string) Worker {
    phases := loadPhases(configfile)
    cycle := makeCycle(phases)
    return Worker{cycle}
}
