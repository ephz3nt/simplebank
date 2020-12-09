package main

import (
	"fmt"
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {

	k := len(alphabet)
	fmt.Println(string(alphabet[rand.Intn(k)]))
}
