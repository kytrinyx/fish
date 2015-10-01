// Command fish embellishes a message.
//
// If the msg flag is not passed, then "Hello, World!" is used as the default message.
package main

import (
	"flag"
	"fmt"
)

var msg = flag.String("msg", "Hello, World!", "What the will say.")

func main() {
	flag.Parse()

	fmt.Println(NewFish(nil).Say(*msg))
}
