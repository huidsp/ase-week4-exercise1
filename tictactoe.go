package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func (b *Board) input() {
	fmt.Print("(x,y): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	in := scanner.Text()
	if x, err := strconv.Atoi(in[0:1]); err == nil {
		if y, err := strconv.Atoi(in[2:3]); err == nil {
			b.put(x, y, "x")
		} else {
			panic(in)
		}
	} else {
		panic(in)
	}
}

func main() {
	b := NewBoard()
	b.input()
	b.print()
}
