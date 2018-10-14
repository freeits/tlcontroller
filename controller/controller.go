package controller

import (
    "log"
    "sync"
    "time"
)

const LOCAL_MODE = 1

type Controller struct {
    cycle Cycle
}

func (m Controller) Run(wg *sync.WaitGroup, ch chan int) {
    defer wg.Done()
    log.Print("Controller started")
    for true {
        phase := m.cycle.nextPhase()
        phase_time := time.Duration(phase.time.total_seconds) * time.Second
        time.Sleep(phase_time)
        log.Printf("Phase %v is active", phase.name)
        ch <- phase.id
        ch <- LOCAL_MODE
    }
}

func MakeController() Controller {
    xmlPhases := loadPhasesXml()
    cycle := makeCycle(xmlPhases)
    return Controller{cycle}
}
