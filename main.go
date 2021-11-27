package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/janakerman/incident-monte-carlo/incident"
	log "github.com/sirupsen/logrus"
)

const (
	// Monte Carlo
	trials = 10_000_000

	// Time frame
	timeStart = 0
	timeEnd   = 60 * 24 * 30

	// Systems
	incidentDuration = 60
	numIncidentsA    = 10
	numIncidentsB    = 1
)

type params struct {
	numIncidentsB    int
	numIncidentsA    int
	incidentDuration float64
}

func init() {
	log.SetLevel(log.InfoLevel)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	params := params{
		incidentDuration: incidentDuration,
		numIncidentsB:    numIncidentsB,
		numIncidentsA:    numIncidentsA,
	}

	log.Infof("Parameters: %+v", params)

	successes := 0

	for i := 0; i < trials; i++ {

		success := simulate(params)
		if success {
			successes++
		}
	}

	probability := float64(successes) / float64(trials)
	log.Infof("Probability: %f", probability)
}

func simulate(p params) bool {
	aIncidents := nonOverlappingIncidents(p.numIncidentsA)
	bIncidents := nonOverlappingIncidents(p.numIncidentsB)

	log.Debugf("A incidents: %v", aIncidents)
	log.Debugf("B incidents: %v", bIncidents)

	anyIncidentsOverlap := anyIncidentsOverlap(aIncidents, bIncidents)
	log.Debugf("Overlap: %v", anyIncidentsOverlap)

	return !anyIncidentsOverlap
}

func bNumIncidents(availabilityB float64, incidentDuration float64) int {
	return int(math.Ceil(float64((timeEnd-timeStart)*(1-availabilityB)) / incidentDuration))
}

func nonOverlappingIncidents(num int) []incident.Incident {
	var incidents []incident.Incident
	for n := 0; n < num; n++ {
		i := incident.NewRandMutex(timeStart, timeEnd, incidentDuration, incidents)
		incidents = append(incidents, i)
	}
	return incidents
}

func anyIncidentsOverlap(is1, is2 []incident.Incident) bool {
	overlap := false
	for _, i1 := range is1 {
		for _, i2 := range is2 {
			overlap = overlap || i1.Overlap(i2)
		}
	}
	return overlap
}
