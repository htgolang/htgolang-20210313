package math

import (
	"fmt"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	// 输入和预期的结果
	if Add(1, 2) != 3 {
		t.Error("1+2 != 3")
	}

}

func TestSub(t *testing.T) {
	if Sub(1, 2) != -1 {
		t.Error("1-2 != -1")
	}
}

func TestMult(t *testing.T) {
	if Mult(0, 10) != 0 {
		t.Error("0 * 10 != 0")
	}
}

func TestFact(t *testing.T) {
	t.Run("fact(0)", func(t *testing.T) {
		if Fact(0) != 1 {
			t.Error("fact(0) != 1")
		}
	})

	t.Run("fact(1)", func(t *testing.T) {
		if Fact(1) != 1 {
			t.Error("fact(1) != 1")
		}
	})

	t.Run("fact(-1)", func(t *testing.T) {
		if Fact(-1) != -1 {
			t.Error("fact(-1) != -1")
		}
	})

}

// setup, teardown
func TestMain(m *testing.M) {
	// setup
	fmt.Println("start")
	code := m.Run()
	fmt.Println("end")
	// teardown
	os.Exit(code)
}
