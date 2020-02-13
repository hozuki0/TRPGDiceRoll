package diceroll

import (
	"testing"
)

func TestDiceRoll(t *testing.T) {
	result, err := DiceRoll("1d6")
	if err == nil {
		t.Errorf("DiceRoll returns nil")
	}
	if !isInRange(result, 1, 6) {
		t.Errorf("DiceRoll result is out of range")
	}
}

func TestExtractDiceSyntax(t *testing.T) {
	ans := "1000d600000"
	result := extractDiceSyntax(ans + "aaaaaaaaaaaaaaaa")
	if result != ans {
		t.Errorf("expected " + ans + " but returns " + result)
	}
	ans = "1d6"
	result = extractDiceSyntax("1D6aaaaaaaaaaaaaaaa")
	if result != ans {
		t.Errorf("expected " + ans + " but returns " + result)
	}
	result = extractDiceSyntax("aaaaaaaaaaaaaaaa")
	if result != "" {
		t.Errorf("expected empty but returns anything")
	}
}

func TestExtractOperator(t *testing.T) {
	ans := "+10000"
	result := extractOperatorFomular("aiueo" + ans)
	if result != ans {
		t.Errorf("expected " + ans + " but returns " + result)
	}

	ans = "*10"
	result = extractOperatorFomular("aiueo " + ans)
	if result != ans {
		t.Errorf("expected " + ans + " but returns " + result)
	}

	result = extractOperatorFomular("aaaaaaaaaaaaaaaa")
	if result != "" {
		t.Errorf("expected empty but returns anything")
	}
}

func isInRange(value int, min int, max int) bool {
	if value < min {
		return false
	}
	if value > max {
		return false
	}
	return true
}
