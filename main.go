// Command fish embellishes a message.
//
// If the msg flag is not passed, then "Hello, World!" is used as the default message.
package main

import (
	"flag"
	"fmt"
	"math/rand"

	"github.com/kytrinyx/fish/fish"
)

var msg = flag.String("msg", "Hello, World!", "What the will say.")

func main() {
	flag.Parse()

	fmt.Println(Doodle(*msg, nil))
}

// Doodle makes random ascii art of a talking fish.
func Doodle(msg string, prng *rand.Rand) string {
	return fish.New(prng).Say(msg)
}
