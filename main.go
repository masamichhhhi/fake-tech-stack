package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/tjarratt/babble"
)

func main() {
	babbler := babble.NewBabbler()
	babbler.Count = 1

	stacks := []string{}

	for {
		if len(stacks) >= 10 {
			break
		}

		if rand.Intn(2) == 1 {
			stacks = append(stacks, RandomString(rand.Intn(5)+2))
		} else {
			for {
				word := babbler.Babble()

				if len(word) <= 6 {
					stacks = append(stacks, strings.Title(word))
					break
				}
			}
		}
	}

	fmt.Println(strings.Join(stacks, "/"))

}

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	if rand.Intn(2) == 1 {
		return strings.ToUpper(string(b))

	} else {
		return strings.Title(string(b))
	}
}
