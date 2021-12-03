package utils

import (
	"fmt"
	"testing"
)

func TestDealIntoNewStack(t *testing.T) {

	pack := NewPack(10)
	pack.DealIntoNewStack()

	expect := "9,8,7,6,5,4,3,2,1,0"
	got := pack.LimitedCSV(10)
	if got != expect {
		t.Errorf("expected %s after DealIntoNewStack. Got %s", expect, got)
	}
}

func TestCut(t *testing.T) {

	pack := NewPack(10)
	pack.Cut(3)

	expect := "3,4,5,6,7,8,9,0,1,2"
	got := pack.LimitedCSV(10)
	if got != expect {
		t.Errorf("expected %s after Cut:3. Got %s", expect, got)
	}
}

func TestCutNegative(t *testing.T) {

	pack := NewPack(10)
	pack.Cut(-4)

	expect := "6,7,8,9,0,1,2,3,4,5"
	got := pack.LimitedCSV(10)
	if got != expect {
		t.Errorf("expected %s after Cut:-4. Got %s", expect, got)
	}
}

func TestDealWithIncrement(t *testing.T) {

	pack := NewPack(10)
	pack.DealWithIncrement(3)

	expect := "0,7,4,1,8,5,2,9,6,3"
	got := pack.LimitedCSV(10)
	if got != expect {
		t.Errorf("expected %s after DealWithIncrement:3. Got %s", expect, got)
	}
}

func TestDealWithIncrementTwice(t *testing.T) {

	pack := NewPack(10)
	pack.DealWithIncrement(3)
	pack.DealWithIncrement(3)

	expect := "0,9,8,7,6,5,4,3,2,1"
	got := pack.LimitedCSV(10)
	if got != expect {
		t.Errorf("expected %s after DealWithIncrementTwice:3. Got %s", expect, got)
	}
}

func TestDealWithIncrementNMinus1(t *testing.T) {

	pack := NewPack(10)
	pack.DealWithIncrement(9)

	expect := "0,9,8,7,6,5,4,3,2,1"
	got := pack.LimitedCSV(10)
	if got != expect {
		t.Errorf("expected %s after TestDealWithIncrementNMinus1:9. Got %s", expect, got)
	}
}

func TestDay22Example1(t *testing.T) {
	pack := NewPack(10)
	pack.DealWithIncrement(7)
	pack.DealIntoNewStack()
	pack.DealIntoNewStack()

	expect := "0,3,6,9,2,5,8,1,4,7"
	got := pack.LimitedCSV(10)
	if got != expect {
		t.Errorf("expected %s after day22 example 1. Got %s", expect, got)
	}
}

func TestDay22Example2(t *testing.T) {
	pack := NewPack(10)
	pack.Cut(6)
	pack.DealWithIncrement(7)
	pack.DealIntoNewStack()

	expect := "3,0,7,4,1,8,5,2,9,6"
	got := pack.LimitedCSV(10)
	if got != expect {
		t.Errorf("expected %s after day22 example 2. Got %s", expect, got)
	}
}
func TestDay22Example3a(t *testing.T) {
	pack := NewPack(10)
	pack.DealWithIncrement(7)

	expect := "0,3,6,9,2,5,8,1,4,7"
	got := pack.LimitedCSV(10)
	if got != expect {
		t.Errorf("expected %s after day22 example 3a. Got %s", expect, got)
	}
}

func TestDay22Example3b(t *testing.T) {
	pack := NewPack(10)
	pack.DealWithIncrement(7)
	pack.DealWithIncrement(9)

	expect := "0,7,4,1,8,5,2,9,6,3"
	got := pack.LimitedCSV(10)
	if got != expect {
		t.Errorf("expected %s after day22 example 3b. Got %s", expect, got)
	}
}

func TestDay22Example3c(t *testing.T) {
	pack := NewPack(10)
	pack.DealWithIncrement(7)
	pack.DealWithIncrement(9)
	pack.Cut(-2)

	expect := "6,3,0,7,4,1,8,5,2,9"
	got := pack.LimitedCSV(10)
	if got != expect {
		t.Errorf("expected %s after day22 example 3c. Got %s", expect, got)
	}
}
func TestDay22Example3cc(t *testing.T) {
	pack := NewPack(10)
	pack.DealWithIncrement(7)
	pack.DealWithIncrement(9)
	pack.DealIntoNewStack()

	expect := "3,6,9,2,5,8,1,4,7,0"
	got := pack.LimitedCSV(10)
	if got != expect {
		t.Errorf("expected %s after day22 example 3cc. Got %s", expect, got)
	}
}

func TestDay22Example4(t *testing.T) {
	pack := NewPack(10)
	pack.DealIntoNewStack()
	pack.Cut(-2)
	pack.DealWithIncrement(7)
	pack.Cut(8)
	pack.Cut(-4)
	pack.DealWithIncrement(7)
	pack.Cut(3)
	pack.DealWithIncrement(9)
	pack.DealWithIncrement(3)
	pack.Cut(-1)

	expect := "9,2,5,8,1,4,7,0,3,6"
	got := pack.LimitedCSV(10)
	if got != expect {
		t.Errorf("expected %s after day22 example 4. Got %s", expect, got)
	}
}

func TestReapplyCuts(t *testing.T) {
	doCut := func(p SpaceCards) {
		p.Cut(2)
		p.Cut(-3)
		p.Cut(-6)
	}
	pack := NewPack(10)
	doCut(pack)
	pack.Reapply(5)
	reapplied := pack.LimitedCSV(10)

	pack = NewPack(10)
	doCut(pack)
	doCut(pack)
	doCut(pack)
	doCut(pack)
	doCut(pack)
	doCut(pack)
	manualReapply := pack.LimitedCSV(10)
	fmt.Println(manualReapply)
	if reapplied != manualReapply {
		t.Errorf("expected reapply of cut sequence to match manual reapplication. Expected %s, got %s", manualReapply, reapplied)
	}
}

func TestReapplyDealWithIncrement(t *testing.T) {
	pack := NewPack(10)
	pack.DealWithIncrement(3)
	pack.Reapply(3)
	reapplied := pack.LimitedCSV(10)

	pack = NewPack(10)
	pack.DealWithIncrement(3)
	pack.DealWithIncrement(3)
	pack.DealWithIncrement(3)
	pack.DealWithIncrement(3)
	manualReapply := pack.LimitedCSV(10)
	fmt.Println(manualReapply)
	if reapplied != manualReapply {
		t.Errorf("expected reapply of deal with increment to match manual reapplication. Expected %s, got %s", manualReapply, reapplied)
	}
}

func TestReapplyDealIntoNewStack(t *testing.T) {
	pack := NewPack(10)
	pack.DealIntoNewStack()
	pack.Reapply(3)
	reapplied := pack.LimitedCSV(10)

	pack = NewPack(10)
	pack.DealIntoNewStack()
	pack.DealIntoNewStack()
	pack.DealIntoNewStack()
	pack.DealIntoNewStack()
	manualReapply := pack.LimitedCSV(10)
	fmt.Println(manualReapply)
	if reapplied != manualReapply {
		t.Errorf("expected reapply of deal into new stack to match manual reapplication. Expected %s, got %s", manualReapply, reapplied)
	}
}
