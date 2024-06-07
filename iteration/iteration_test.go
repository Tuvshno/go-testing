package iteration

import "testing"

func TestRepeat(t *testing.T) {
	get := Repeat("a")
	expect := "aaaaa"

	if get != expect {
		t.Errorf("wanted %q got %q", get, expect)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
