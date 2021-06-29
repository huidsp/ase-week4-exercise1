package main

import "fmt"

type Board struct {
	tokens []string
}

func NewBoard() *Board {
	return &Board{
		tokens: make([]string, 9),
	}
}

func (b *Board) put(x, y int, v string) {
	b.tokens[y+3*x] = v

}

func (b *Board) get(x, y int) string {
	return b.tokens[y+3*x]
}

func (b *Board) print() {
	for i, x := range b.tokens {
		if x == "o" || x == "x" {
			fmt.Print(x)
		} else {
			fmt.Print(".")
		}
		if i%3 == 2 {
			fmt.Println()
		}
	}
}
