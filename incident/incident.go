package incident

import "math/rand"

type Incidents []Incident

func (is Incidents) Overlap(is2 Incidents) bool {
	overlap := false
	for _, i1 := range is {
		for _, i2 := range is2 {
			overlap = overlap || i1.Overlap(i2)
		}
	}
	return overlap
}

type Incident [2]int

func NewRand(periodStart, periodEnd, incidentDuration int) Incident {
	start := rand.Intn(periodEnd-periodStart) + periodStart
	return Incident{start, start + incidentDuration}
}

func NewRandMutex(start, end, incidentDuration int, mutexIncidents []Incident) Incident {
	var incident Incident
	for {
		potential := NewRand(start, end, incidentDuration)
		overlap := false
		for _, i2 := range mutexIncidents {
			if potential.Overlap(i2) {
				overlap = true
			}
		}
		if !overlap {
			incident = potential
			break
		}
	}

	return incident
}

func (i Incident) Start() int {
	return i[0]
}

func (i Incident) End() int {
	return i[1]
}

func (i Incident) Overlap(i2 Incident) bool {
	return i.Start() >= i2.Start() && i.Start() < i2.End() ||
		i.End() > i2.Start() && i.End() <= i2.End()
}
