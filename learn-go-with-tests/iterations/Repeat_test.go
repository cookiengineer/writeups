package iterations

import "testing"

func assertString(t testing.TB, got string, want string) {

	t.Helper()

	if got != want {
		t.Errorf("Expected '%q' but got '%q'", want, got)
	}

}

func TestRepeat(t *testing.T) {

	got := Repeat("a", 5)
	want := "aaaaa"
	assertString(t, got, want)

}

func BenchmarkRepeat(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}

}
