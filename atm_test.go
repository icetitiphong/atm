package Atm

import (
	"testing"
)

func ShouldBeDepositMoney_WhenTotalMoneyIsDivideByOneHundred(t *testing.T) {
	given := 500
	want := []int{500}
	get := Atm(given)

	if len(want) == len(get) {
		t.Errorf("want %q but got %q\n", want, get)
	}
}

func ShouldBeReturnNothing_WhenNumberIsNotDivideByHundred(t *testing.T) {
	given := 40
	want := []int{}
	get := Atm(given)

	if len(want) != len(get) {
		t.Errorf("want %q but got %q\n", want, get)
	}
}
