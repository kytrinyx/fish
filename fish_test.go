package main

import (
	"math/rand"
	"testing"
)

const errMsg = `%s looks weird (seed: %d).

Got:
%s

Want:
%s

`

// Not "real" unit tests, per se. I wrote them
// as golden master tests, and am leaving them
// in place because they work well as a regression.
func TestFishBody(t *testing.T) {
	tests := []struct {
		seed int64
		body string
		said string
	}{
		{2, "><<˚>", "         ºº. OHAI o.•.\n><<˚> ·°"},
		{3, "><(((˚>", "           •˚˚ OHAI .º.˚\n><(((˚> ·°"},
		{5, "><<<º>", "          ˚˚o OHAI oº˚.\n><<<º> ·°"},
		{7, "><<˚>", "         ˚ºº OHAI ˚o˚.\n><<˚> ·°"},
		{11, "><((('>", "          ˚.. OHAI .O˚•\n><((('> ·°"},
		{13, "><(˚>", "         ˚.• OHAI ˚O˚.\n><(˚> ·°"},
		{17, "><{{{º>", "           º˚. OHAI ˚º..\n><{{{º> ·°"},
	}

	for _, test := range tests {
		f := NewFish(rand.New(rand.NewSource(test.seed)))
		if f.body != test.body {
			t.Errorf(errMsg, "Fish", test.seed, f.body, test.body)
		}

		said := f.Say("OHAI")
		if said != test.said {
			t.Errorf(errMsg, "Speech", test.seed, said, test.said)
		}
	}
}
