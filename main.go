// Command fish embellishes a message.
//
// If the msg flag is not passed, then "Hello, World!" is used as the default message.
package main

import (
	"flag"
	"fmt"

	"github.com/kytrinyx/fish/bubbles"
)

var msg = flag.String("msg", "Hello, World!", "The text that you want embellished.")

func main() {
	flag.Parse()

	fmt.Println(bubbles.Embellish(*msg))
}
