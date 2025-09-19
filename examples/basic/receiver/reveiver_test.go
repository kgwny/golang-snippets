package main

import "testing"

func TestArea(t *testing.T) {
	tests := []struct {
		name     string
		rect     Rectangle
		expected int
	}{
		{"2x3", Rectangle{2, 3}, 6},
		{"4x5", Rectangle{4, 5}, 20},
		{"0x10", Rectangle{0, 10}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rect.Area(); got != tt.expected {
				t.Errorf("Area() = %d, want %d", got, tt.expected)
			}
		})
	}
}

func TestScale(t *testing.T) {
	rect := Rectangle{2, 3}
	rect.Scale(2)

	if rect.Width != 4 || rect.Height != 6 {
		t.Errorf("Scale() failed: got = {%d %d}, want = {4 6}", rect.Width, rect.Height)
	}

	// 複数回スケールした場合の確認
	rect.Scale(3) // {4*3, 6*3} = {12, 18}
	if rect.Width != 12 || rect.Height != 18 {
		t.Errorf("Scale() failed after second scale: got = {%d %d}, want = {12 18}", rect.Width, rect.Height)
	}
}
