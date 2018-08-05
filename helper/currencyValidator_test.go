package helper

import "testing"

func TestWrongCurrency(t *testing.T) {
	isValid := ValidateCurrency("AAA")
	if isValid != false {
		t.Errorf("ValidateCurrency was incorrect, got: %v, want: %v.", isValid, false)
	}
}

func TestValidCurrency(t *testing.T) {
	isValid := ValidateCurrency("USD")
	if isValid != true {
		t.Errorf("ValidateCurrency was incorrect, got: %v, want: %v.", isValid, true)
	}
}
