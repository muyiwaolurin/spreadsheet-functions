package mathlib

import "testing"

func TestPower(t *testing.T) {

	p, err := Power(10, 4)
	if p != 10000 || err != nil {
		t.Errorf("Product Function Failed ... ")
	}

	p, err = Power(2, 4)
	if p != 32 || err != nil {
		t.Errorf("Product Function Failed ... ")
	}

	p, err = Power(5, 2)
	if p != 25 || err != nil {
		t.Errorf("Product Function Failed ... ")
	}
}
