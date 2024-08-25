package maps

const (
	ErrNotFound = DictionaryError("Cannot find Word in Dictionary")
	ErrWordExists = DictionaryError("Cannot add Word because it already exists in Dictionary")
	ErrWordDoesNotExist = DictionaryError("Cannot update Word because it does not exist in Dictionary")
)

type DictionaryError string

func (err DictionaryError) Error() string {
	return string(err)
}

type Dictionary map[string]string

func (dict Dictionary) Search(word string) (string, error) {

	definition, ok := dict[word]

	if ok == true {
		return definition, nil
	}

	return "", ErrNotFound

}

func (dict Dictionary) Add(word, definition string) error {

	_, err := dict.Search(word)

	switch err {
	case ErrNotFound:
		dict[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil

}

func (dict Dictionary) Update(word, definition string) error {

	_, err := dict.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		dict[word] = definition
	default:
		return err

	}

	return nil

}

func (dict Dictionary) Delete(word string) {
	delete(dict, word)
}
