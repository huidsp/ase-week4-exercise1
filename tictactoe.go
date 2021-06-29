package main

import "fmt"

type Board struct {
	tokens string
}

func NewBoard() *Board {
	return &Board{""}
}

func (b *Board) put(x, y int, v string) {
	b.tokens = v

}

func (b *Board) get(x, y int) string {
	return b.tokens
}

func (b *Board) print() {
	fmt.Println(b.tokens)
}
