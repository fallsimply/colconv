package colconv

import (
	"testing"
)

// Color A = hsl(156, 50%, 50%) = rgb(64, 191, 140)
// Color B = hsl(291, 25%, 25%) = rgb(75, 48, 80)
// Color C = hsl(0, 0%, 80%)    = rgb(204, 204, 204)
// Color D = hsl(0, 40%, 40%)   = rgb(143, 61, 61)

func TestRgbToHsl(t *testing.T) {
	strA := RgbToHsl(64, 191, 140)
	strB := RgbToHsl(75, 48, 80)
	strC := RgbToHsl(204, 204, 204)
	strD := RgbToHsl(143, 61, 61)
	switch {
	case strA != `hsl(156, 50%, 50%)`:
		t.Error(strA)
	case strB != `hsl(291, 25%, 25%)`:
		t.Error(strB)
	case strC != `hsl(0, 0%, 80%)`:
		t.Error(strC)
	case strD != `hsl(0, 40%, 40%)`:
		t.Error(strD)
	}
}
func BenchmarkRgbToHsl(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = RgbToHsl(uint8(i), uint8(255-i), uint8(i)) + " "
	}
}
