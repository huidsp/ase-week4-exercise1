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
	fmt.Println("x..")
	fmt.Println("...")
	fmt.Println("...")
	b.put(0, 0, "x")
	b.print()
}

func TestPut04(t *testing.T) {
	b := NewBoard()
	fmt.Println("expected:")
	fmt.Println("o..")
	fmt.Println("...")
	fmt.Println("...")
	b.put(0, 0, "o")
	b.print()
}

func TestPut05(t *testing.T) {
	b := NewBoard()
	b.put(0, 1, "o")
	result := b.get(0, 1)
	if result != "o" {
		t.Errorf("expected %v, result %v", "o", result)
	}
}

func TestPut06(t *testing.T) {
	b := NewBoard()
	fmt.Println("expected:")
	fmt.Println(".o.")
	fmt.Println("...")
	fmt.Println("...")
	b.put(0, 1, "o")
	b.print()
}

func TestPut07(t *testing.T) {
	b := NewBoard()
	fmt.Println("expected:")
	fmt.Println("...")
	fmt.Println("x..")
	fmt.Println("...")
	b.put(1, 0, "x")
	b.print()
}

func TestPut08(t *testing.T) {
	b := NewBoard()
	b.put(0, 0, "o")
	b.put(1, 1, "o")
	b.put(2, 2, "o")
	result := b.check()
	if result != "o" {
		t.Errorf("expected %v, result %v", "o", result)
	}
}

func TestPut09(t *testing.T) {
	b := NewBoard()
	b.put(0, 2, "o")
	b.put(1, 1, "o")
	b.put(2, 0, "o")
	result := b.check()
	if result != "o" {
		t.Errorf("expected %v, result %v", "o", result)
	}
}

func TestPut10(t *testing.T) {
	b := NewBoard()
	b.put(0, 0, "o")
	b.put(1, 0, "o")
	b.put(2, 0, "o")
	result := b.check()
	if result != "o" {
		t.Errorf("expected %v, result %v", "o", result)
	}
}

func TestPut11(t *testing.T) {
	b := NewBoard()
	b.put(0, 1, "o")
	b.put(1, 1, "o")
	b.put(2, 1, "o")
	result := b.check()
	if result != "o" {
		t.Errorf("expected %v, result %v", "o", result)
	}
}

func TestPut12(t *testing.T) {
	b := NewBoard()
	b.put(0, 2, "o")
	b.put(1, 2, "o")
	b.put(2, 2, "o")
	result := b.check()
	if result != "o" {
		t.Errorf("expected %v, result %v", "o", result)
	}
}

func TestPut13(t *testing.T) {
	b := NewBoard()
	b.put(0, 0, "o")
	b.put(0, 1, "o")
	b.put(0, 2, "o")
	result := b.check()
	if result != "o" {
		t.Errorf("expected %v, result %v", "o", result)
	}
}

func TestPut14(t *testing.T) {
	b := NewBoard()
	b.put(1, 0, "o")
	b.put(1, 1, "o")
	b.put(1, 2, "o")
	result := b.check()
	if result != "o" {
		t.Errorf("expected %v, result %v", "o", result)
	}
}

func TestPut15(t *testing.T) {
	b := NewBoard()
	b.put(2, 0, "o")
	b.put(2, 1, "o")
	b.put(2, 2, "o")
	result := b.check()
	if result != "o" {
		t.Errorf("expected %v, result %v", "o", result)
	}
}

func TestPut16(t *testing.T) {
	b := NewBoard()
	b.put(0, 0, "x")
	b.put(1, 1, "x")
	b.put(2, 2, "x")
	result := b.check()
	if result != "x" {
		t.Errorf("expected %v, result %v", "x", result)
	}
}

func TestPut17(t *testing.T) {
	b := NewBoard()
	b.put(0, 2, "x")
	b.put(1, 1, "x")
	b.put(2, 0, "x")
	result := b.check()
	if result != "x" {
		t.Errorf("expected %v, result %v", "x", result)
	}
}

func TestPut18(t *testing.T) {
	b := NewBoard()
	b.put(0, 0, "x")
	b.put(1, 0, "x")
	b.put(2, 0, "x")
	result := b.check()
	if result != "x" {
		t.Errorf("expected %v, result %v", "x", result)
	}
}

func TestPut19(t *testing.T) {
	b := NewBoard()
	b.put(0, 1, "x")
	b.put(1, 1, "x")
	b.put(2, 1, "x")
	result := b.check()
	if result != "x" {
		t.Errorf("expected %v, result %v", "x", result)
	}
}

func TestPut20(t *testing.T) {
	b := NewBoard()
	b.put(0, 2, "x")
	b.put(1, 2, "x")
	b.put(2, 2, "x")
	result := b.check()
	if result != "x" {
		t.Errorf("expected %v, result %v", "x", result)
	}
}

func TestPut21(t *testing.T) {
	b := NewBoard()
	b.put(0, 0, "x")
	b.put(0, 1, "x")
	b.put(0, 2, "x")
	result := b.check()
	if result != "x" {
		t.Errorf("expected %v, result %v", "x", result)
	}
}

func TestPut22(t *testing.T) {
	b := NewBoard()
	b.put(1, 0, "x")
	b.put(1, 1, "x")
	b.put(1, 2, "x")
	result := b.check()
	if result != "x" {
		t.Errorf("expected %v, result %v", "x", result)
	}
}

func TestPut23(t *testing.T) {
	b := NewBoard()
	b.put(2, 0, "x")
	b.put(2, 1, "x")
	b.put(2, 2, "x")
	result := b.check()
	if result != "x" {
		t.Errorf("expected %v, result %v", "x", result)
	}
}

func TestPut24(t *testing.T) {
	b := NewBoard()
	b.put(0, 0, "o")
	b.put(0, 1, "x")
	b.put(0, 2, "o")
	b.put(1, 0, "x")
	b.put(1, 1, "o")
	b.put(1, 2, "x")
	b.put(2, 0, "x")
	b.put(2, 1, "o")
	b.put(2, 2, "x")
	result := b.check()
	if result != "d" {
		t.Errorf("expected %v, result %v", "d", result)
	}
}

func TestPut25(t *testing.T) {
	b := NewBoard()
	b.put(0, 0, "o")
	b.put(0, 1, "x")
	b.put(1, 0, "x")
	b.put(1, 1, "o")
	b.put(1, 2, "x")
	b.put(2, 0, "x")
	b.put(2, 1, "o")
	b.put(2, 2, "x")
	result := b.check()
	if result != "" {
		t.Errorf("expected %v, result %v", "", result)
	}
}
