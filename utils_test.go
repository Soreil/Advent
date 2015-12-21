package advent

import "testing"

func TestPermutationsString(t *testing.T) {
	t.Log(PermutationsString([]string{"ayy", "may", "bay"}))
}

func TestPermutationsInt(t *testing.T) {
	t.Log(PermutationsInt([]int{1, 2, 3, 4, 5}))
}
