package testing101

import (
  "testing"
)
 
func TestSum1(t *testing.T) {
  numbers := []int{1, 2, 3, 4, 5}
  expected := 15
  actual := Sum(numbers)

  if actual != expected {
    t.Errorf("Expected the sum of %v to be %d but instead got %d!", numbers, expected, actual)
  }
}