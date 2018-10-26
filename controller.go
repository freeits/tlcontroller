package controller

import "sync"

func Run(configfile string, addr string){
    worker := makeWorker(configfile)
    server := makeServer(addr)

    var wg sync.WaitGroup
    wg.Add(2)

    state := NewControllerState()
    go worker.work(&wg, state)
    go server.serve(&wg, state)
    wg.Wait()
}
