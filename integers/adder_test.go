package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("Test Adder", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected : '%d' but got '%d'", expected, sum)
		}
	})

}
