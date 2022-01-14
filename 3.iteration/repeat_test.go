package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expacted := "aaaaa"

	if repeated != expacted {
		t.Errorf("expacted %q but got %q", expacted, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
