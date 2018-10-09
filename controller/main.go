package main
import (
    "fmt"
)

type PhaseTime struct {
    total_seconds int
    min_seconds int
    ending_seconds int
}

type Phase struct {
    id int
    name string
    time PhaseTime
    next_phase *Phase
}

type Cycle struct {
    current_phase *Phase
    phases []Phase
}

func (c *Cycle) next_phase() *Phase {
    c.current_phase = c.current_phase.next_phase
    return c.current_phase
}


func main(){

    phase3 := Phase{3, "Phase 3", PhaseTime{7, 3, 3}, nil}
    phase2 := Phase{2, "Phase 2", PhaseTime{12, 4, 4}, &phase3}
    phase1 := Phase{1, "Phase 1", PhaseTime{10, 3, 3}, &phase2}
    phase3.next_phase = &phase1

    cycle := Cycle{&phase1, []Phase{phase1, phase2, phase3}}
    cycle.current_phase = &phase1

    for i := 0; i < 10; i++ {
        phase := cycle.next_phase()
        fmt.Println(phase.name)
    }

}
