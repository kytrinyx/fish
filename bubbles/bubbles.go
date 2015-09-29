package bubbles

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	oo   = []string{".", ".", ".", "o", "o", "O"}
	prng = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Embellish decorates the string with ascii bubbles.
func Embellish(s string) string {
	return fmt.Sprintf("%s %s %s", Blub(3), s, Blub(4))
}

// Blub concatenates n randomly-sized bubbles.
func Blub(n int) string {
	sx := make([]string, n)
	for i := 0; i < n; i++ {
		sx[i] = oo[rand.Intn(len(oo))]
	}
	return strings.Join(sx, "")
}
