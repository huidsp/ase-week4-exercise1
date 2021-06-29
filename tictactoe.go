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

func (b *Board) check_o1(x string) string {
	for i := 0; i < 3; i++ {
		if b.get(i, i) != x {
			return ""
		}
	}
	return x
}

func (b *Board) check_o2(x string) string {
	for i := 0; i < 3; i++ {
		if b.get(i, 2-i) != x {
			return ""
		}
	}
	return x
}

func (b *Board) check_o3(j int, x string) string {
	for i := 0; i < 3; i++ {
		if b.get(i, j) != x {
			return ""
		}
	}
	return x
}

func (b *Board) check_o4(j int, x string) string {
	for i := 0; i < 3; i++ {
		if b.get(j, i) != x {
			return ""
		}
	}
	return x
}

func (b *Board) check_o5() string {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.get(j, i) == "" {
				return ""
			}
		}
	}
	return "d"
}

func (b *Board) check() string {
	for _, x := range []string{"o", "x"} {
		if b.check_o1(x) == x || b.check_o2(x) == x {
			return x
		}
		for j := 0; j < 3; j++ {
			if b.check_o3(j, x) == x {
				return x
			}
			if b.check_o4(j, x) == x {
				return x
			}
		}
	}
	return b.check_o5()
}

func main() {
	b := NewBoard()
	b.input()
	b.print()
}
