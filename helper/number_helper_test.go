package helper

import "testing"

func TestRound(t *testing.T) {
	tests := []struct {
		name     string
		in       float64
		expected int
	}{
		{"ปัดลง", 3.2, 3},
		{"ครึ่งควรปัดขึ้น", 2.5, 3},
		{"จำนวนเต็ม", 5.0, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := Round(tt.in); result != tt.expected {
				t.Errorf("Rounded(%v) = %d; want %d", tt.in, result, tt.expected)
			}
		})
	}
}

func TestToFixed(t *testing.T) {
	basenum := 2.3456789
	tests := []struct {
		name      string
		precision int
		expected  float64
	}{
		{"toFixed Case 1", 1, 2.3},
		{"toFixed Case 2", 2, 2.34},
		{"toFixed Case 3", 3, 2.345},
		{"toFixed Case 4", 4, 2.3456},
		{"toFixed Case 5", 5, 2.34567},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := toFixed(basenum, tt.precision); result != tt.expected {
				t.Errorf("toFixed(%f) = %f; want %f", basenum, result, tt.expected)
			}
		})
	}
}
