package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	fmt.Println(randomPalindrome(rng))
}
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}
