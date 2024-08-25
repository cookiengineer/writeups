package maps

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is a test"
		assertStrings(t, got, want)

	})

	t.Run("unknown word", func(t *testing.T) {

		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)

	})

}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {

		dictionary := Dictionary{}
		word := "test"
		definition := "this is a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)

	})

	t.Run("existing word", func(t *testing.T) {

		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)

	})
}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {

		word := "test"
		definition := "this is a test"
		new_definition := "new definition"
		dictionary := Dictionary{word: definition}
		err := dictionary.Update(word, new_definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, new_definition)

	})

	t.Run("new word", func(t *testing.T) {

		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)

	})
}

func TestDelete(t *testing.T) {

	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	assertError(t, err, ErrNotFound)

}

func assertStrings(t testing.TB, got string, want string) {

	t.Helper()

	if got != want {
		t.Errorf("Expected %q and got %q", want, got)
	}

}

func assertError(t testing.TB, got error, want error) {

	t.Helper()

	if got != want {
		t.Errorf("Expected %q and got %q", want, got)
	}

}

func assertDefinition(t testing.TB, dict Dictionary, word string, definition string) {

	t.Helper()

	got, err := dict.Search(word)

	if err != nil {
		t.Fatal("Should find added word:", err)
	}

	if definition != got {
		t.Errorf("Expected %q and got %q", definition, got)
	}

}
