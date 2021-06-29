package main

import (
	"fmt"
	"testing"
)

func TestPut01(t *testing.T) {
	b := NewBoard()
	b.put(0, 0, "o")
	result := b.get(0, 0)
	if result != "o" {
		t.Errorf("expected %v, result %v", "o", result)
	}
}

func TestPut02(t *testing.T) {
	b := NewBoard()
	b.put(0, 0, "x")
	result := b.get(0, 0)
	if result != "x" {
		t.Errorf("expected %v, result %v", "x", result)
	}
}

func TestPut03(t *testing.T) {
	b := NewBoard()
	fmt.Println("expected:")
	fmt.Println("x")
	b.put(0, 0, "x")
	b.print()
}

func TestPut04(t *testing.T) {
	b := NewBoard()
	fmt.Println("expected:")
	fmt.Println("o")
	b.put(0, 0, "o")
	b.print()
}
