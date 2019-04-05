package core

import "testing"

func init() {
	CreateContext()
}

func TestXPForNextLevel(t *testing.T) {
	input := 113901
	expected := 3399
	actual := GetRequired(input)

	if expected != actual {
		t.Errorf("Expected: %v, got %v", expected, actual)
	}
}

func TestLevelForTotalXP(t *testing.T) {
	input := 113901
	expected := 67
	actual := GetRank(input)

	if expected != actual {
		t.Errorf("Expected: %v, got %v", expected, actual)
	}
}
