package controller

import "sync"

func Run(configfile string, addr string){
    worker := makeWorker(configfile)
    server := makeServer(addr)

    var wg sync.WaitGroup
    wg.Add(2)
    ch := make(chan int)
    go worker.work(&wg, ch)
    go server.serve(&wg, ch)
    wg.Wait()
}
