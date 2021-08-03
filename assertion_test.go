package assertion

import "testing"

func TestEquals(t *testing.T) {
	Equals("123", "123")
	Equals(213, 213)
	Equals(123.5, 123.5)
	Equals('A', 'A')
	Equals([]int{32, 33, 34}, []int{32, 33, 34})
	Equals(map[string]string{"A": "bc", "D": "ab"}, map[string]string{"D": "ab", "A": "bc"})
}
