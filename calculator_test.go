package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		a    float64
		b    float64
		want float64
	}{
		{
			name: "TEST multiplying 2 numbers",
			a:    5,
			b:    2,
			want: 10,
		},
		{
			name: "TEST multiplying 0 by a negative number",
			a:    0,
			b:    -1,
			want: 0,
		},
		{
			name: "TEST multiplying 2 negative numbers",
			a:    -1,
			b:    -1,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculator.Multiply(tt.a, tt.b)
			if tt.want != got {
				t.Errorf("want %f, got %f", tt.want, got)
			}
		})
	}

	// var want float64 = 10
	// got := calculator.Multiply(5, 2)

	// if want != got {
	// 	t.Errorf("want %f, got %f", want, got)
	// }
}

// func TestDivision(t *testing.T) {
// 	t.Parallel()
// 	var want float64 = 2
// 	got := calculator.Division(10, 5)
// 	if want != got {
// 		t.Errorf("want %f, got %f", want, got)
// 	}
// }
