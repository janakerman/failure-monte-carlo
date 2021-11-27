package incident

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Incident_Overlap(t *testing.T) {
	tt := []struct {
		name     string
		i1       Incident
		i2       Incident
		expected bool
	}{
		{
			name:     "no overlap",
			i1:       Incident{1, 2},
			i2:       Incident{4, 6},
			expected: false,
		},
		{
			name:     "no overlap end of i1 flush",
			i1:       Incident{1, 2},
			i2:       Incident{2, 3},
			expected: false,
		},
		{
			name:     "overlap end of i1",
			i1:       Incident{1, 3},
			i2:       Incident{2, 4},
			expected: true,
		},
		{
			name:     "overlap start of i1",
			i1:       Incident{1, 3},
			i2:       Incident{0, 2},
			expected: true,
		},
		{
			name:     "overlap complete",
			i1:       Incident{1, 3},
			i2:       Incident{0, 2},
			expected: true,
		},
		{
			name:     "overlap perfect",
			i1:       Incident{1, 2},
			i2:       Incident{1, 2},
			expected: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.i1.Overlap((tc.i2)))
		})
	}
}

func Test_Incident_NewRandMutex(t *testing.T) {
	rand.Seed(1)
	i1 := Incident{0, 9}
	i2 := NewRandMutex(0, 10, 2, []Incident{i1})
	assert.Equal(t, 9, i2.Start())
	assert.Equal(t, 11, i2.End())
}

func Test_Incidents_Overlap(t *testing.T) {
	tt := []struct {
		name     string
		is1      Incidents
		is2      Incidents
		expected bool
	}{
		{
			is1: Incidents{
				Incident{0, 1},
				Incident{1, 2},
				Incident{2, 3},
			},
			is2: Incidents{
				Incident{3, 4},
				Incident{4, 5},
				Incident{5, 6},
			},
			expected: false,
		},
		{
			is1: Incidents{
				Incident{0, 1},
				Incident{1, 2},
				Incident{2, 3},
			},
			is2: Incidents{
				Incident{1, 2},
				Incident{4, 5},
				Incident{5, 6},
			},
			expected: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.is1.Overlap(tc.is2))
		})
	}
}
