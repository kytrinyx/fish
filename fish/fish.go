package fish

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/kytrinyx/fish/bubbles"
)

var (
	fins = []string{"(", "{", "<", "«"}
	eyes = []string{"˚", "º", "'"}
)

// Fish swim around saying bubbly things.
type Fish struct {
	prng *rand.Rand
	Body string
}

func randomBody(prng *rand.Rand) string {
	return fmt.Sprintf("><%s%s>", randomFin(prng), randomEye(prng))
}

func randomEye(prng *rand.Rand) string {
	return eyes[prng.Intn(len(eyes))]
}

func randomFin(prng *rand.Rand) string {
	return strings.Repeat(fins[prng.Intn(len(fins))], prng.Intn(3)+1)
}

// New spawns a fish.
func New(prng *rand.Rand) *Fish {
	if prng == nil {
		prng = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return &Fish{
		prng: prng,
		Body: randomBody(prng),
	}
}

// Say adds bubbles to whatever the fish is saying.
func (f *Fish) Say(s string) string {
	padding := strings.Repeat(" ", len(f.Body)+3)
	msg := bubbles.New(f.prng).Embellish(s)
	return fmt.Sprintf("%s%s\n%s ·°", padding, msg, f.Body)
}
