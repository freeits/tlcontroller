package controller

func makeCycle(xmlPhases XmlPhases) Cycle {
	phases := []Phase{}
	for _, xmlPhase := range xmlPhases.XmlPhases {
		xmlPhaseTime := PhaseTime{xmlPhase.Time.TotalSeconds,
								  xmlPhase.Time.MinSeconds,
								  xmlPhase.Time.EndingSeconds}
		phase := Phase{xmlPhase.Id, xmlPhase.Name, xmlPhaseTime, nil}
		phases = append(phases, phase)
	}
	// Update links
	for i, _ := range phases {
		if i > 0 {
			phases[i - 1].next_phase = &phases[i]
		}
	}
	phases[len(phases) - 1].next_phase = &phases[0]
	return Cycle{&phases[0], phases}
}

type Cycle struct {
	current_phase *Phase
	phases []Phase
}

func (c *Cycle) nextPhase() *Phase {
	c.current_phase = c.current_phase.next_phase
	return c.current_phase
}
