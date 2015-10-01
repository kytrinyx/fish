/*
Package bubbles provides ascii bubble embellishment.
*/
package bubbles

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var oo = []string{".", ".", ".", "˚", "˚", "˚", "•", "•", "º", "º", "o", "O"}

// Bubbles is an ASCII art bubble generator.
type Bubbles struct {
	prng *rand.Rand
}

// New provides Bubbles.
func New(prng *rand.Rand) *Bubbles {
	if prng == nil {
		prng = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return &Bubbles{prng: prng}
}

// Embellish decorates a string with ASCII bubbles.
func (bb *Bubbles) Embellish(s string) string {
	return fmt.Sprintf("%s %s %s", bb.Blub(3), s, bb.Blub(4))
}

// Blub concatenates n randomly-sized bubbles.
func (bb *Bubbles) Blub(n int) string {
	sx := make([]string, n)
	for i := 0; i < n; i++ {
		sx[i] = bb.randomBubble()
	}
	return strings.Join(sx, "")
}

func (bb *Bubbles) randomBubble() string {
	return oo[bb.prng.Intn(len(oo))]
}
