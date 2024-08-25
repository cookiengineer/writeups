package reflection

import "reflect"
import "testing"

type MockDetails struct {
	City string
	Age int
}

func assertContains(t testing.TB, values []string, search string) {

	t.Helper()

	found := false

	for v := 0; v < len(values); v++ {

		if values[v] == search {
			found = true
			break
		}

	}

	if found == false {
		t.Errorf("Expected %v to contain %q", values, search)
	}

}

func TestWalkProperties(t *testing.T) {

	t.Run("struct with string field", func(t *testing.T) {

		data := struct{
			Name string
		}{
			"Cookie Engineer",
		}

		want := []string{"Cookie Engineer"}
		got := []string{}

		WalkProperties(data, func(value string) {
			got = append(got, value)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}

	})

	t.Run("struct with three string fields", func(t *testing.T) {

		data := struct{
			Name string
			City string
			Country string
		}{
			"Jean-Luc Picard",
			"La Barre",
			"France",
		}

		want := []string{"Jean-Luc Picard", "La Barre", "France"}
		got := []string{}

		WalkProperties(data, func(value string) {
			got = append(got, value)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}

	})

	t.Run("struct with string and int fields", func(t *testing.T) {

		data := struct{
			Name string
			Age  int
		}{
			"Jean-Luc Picard",
			94,
		}

		want := []string{"Jean-Luc Picard"}
		got := []string{}

		WalkProperties(data, func(value string) {
			got = append(got, value)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}

	})

	t.Run("struct with nested struct properties", func(t *testing.T) {

		details := MockDetails{
			"La Barre",
			94,
		}
		data := struct{
			Name string
			Details MockDetails
		}{
			"Jean-Luc Picard",
			details,
		}

		want := []string{"Jean-Luc Picard", "La Barre"}
		got := []string{}

		WalkProperties(data, func(value string) {
			got = append(got, value)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}

	})

	t.Run("struct with pointers", func(t *testing.T) {

		details := MockDetails{
			"La Barre",
			94,
		}
		data := struct{
			Name string
			Details *MockDetails
		}{
			"Jean-Luc Picard",
			&details,
		}

		want := []string{"Jean-Luc Picard", "La Barre"}
		got := []string{}

		WalkProperties(data, func(value string) {
			got = append(got, value)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}

	})

	t.Run("struct with slices", func(t *testing.T) {

		data := struct{
			Name string
			Details []MockDetails
		}{
			"Jean-Luc Picard",
			[]MockDetails{
				MockDetails{"La Barre", 0},
				MockDetails{"San Francisco", 16},
			},
		}

		want := []string{"Jean-Luc Picard", "La Barre", "San Francisco"}
		got := []string{}

		WalkProperties(data, func(value string) {
			got = append(got, value)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}

	})

	t.Run("maps", func(t *testing.T) {

		data := map[string]string{
			"foo": "bar",
			"qux": "doo",
		}

		got := []string{}

		WalkProperties(data, func(value string) {
			got = append(got, value)
		})

		assertContains(t, got, "bar")
		assertContains(t, got, "doo")

	})

	t.Run("channels", func(t *testing.T) {

		channel := make(chan MockDetails)

		go func() {
			channel <- MockDetails{"Berlin", 13}
			channel <- MockDetails{"Munich", 37}
			close(channel)
		}()

		got := []string{}
		want := []string{"Berlin", "Munich"}

		WalkProperties(channel, func(value string) {
			got = append(got, value)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}

	})

	t.Run("with function", func(t *testing.T) {

		function := func() (MockDetails, MockDetails) {
			return MockDetails{"Berlin", 13}, MockDetails{"Munich", 37}
		}
	
		got := []string{}
		want := []string{"Berlin", "Munich"}
	
		WalkProperties(function, func(value string) {
			got = append(got, value)
		})
	
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}

	})

}

