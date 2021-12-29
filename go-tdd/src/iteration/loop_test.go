package iteration

import "testing"

func TestLoopStringFive(t *testing.T) {
	got := Repeat("a")
	expected := "aaaaa"

	if got != expected {
		t.Errorf("expected '%q' but got '%q'", expected, got)
	}
}

// 実行にかかる時間を計測してくれる
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
