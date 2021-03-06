package main

import (
	"math/rand"
	"testing"
)

const errMsg = `Output looks weird (seed: %d).

Got:
%s

Want:
%s

`

// Not "real" unit tests, per se. I wrote them
// as golden master tests, and am leaving them
// in place because they work well as a regression.
func TestDoodle(t *testing.T) {
	tests := []struct {
		seed   int64
		doodle string
	}{
		{2, "         ºº. OHAI o.•.\n><<˚> ·°"},
		{3, "           •˚˚ OHAI .º.˚\n><(((˚> ·°"},
		{5, "          ˚˚o OHAI oº˚.\n><<<º> ·°"},
		{7, "         ˚ºº OHAI ˚o˚.\n><<˚> ·°"},
		{11, "          ˚.. OHAI .O˚•\n><((('> ·°"},
		{13, "         ˚.• OHAI ˚O˚.\n><(˚> ·°"},
		{17, "           º˚. OHAI ˚º..\n><{{{º> ·°"},
	}

	for _, test := range tests {
		doodle := Doodle("OHAI", rand.New(rand.NewSource(test.seed)))
		if doodle != test.doodle {
			t.Errorf(errMsg, test.seed, doodle, test.doodle)
		}
	}
}
