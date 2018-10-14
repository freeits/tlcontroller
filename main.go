package main
import (
    "sync"
    "./controller"
)

func main(){
    ctrl := controller.MakeController()
    var wg sync.WaitGroup
    wg.Add(2)
    ch := make(chan int)
    go ctrl.Run(&wg, ch)
    go controller.Server(&wg, ch)
    wg.Wait()
}
