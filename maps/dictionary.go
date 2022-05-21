package search

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("Item already exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	if err != nil {
		return err
	}
	d[key] = value
	return nil
}

func (d Dictionary) Search(key string) (string, error) {

	value, ok := d[key]
	if !ok {
		return value, ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	if err == ErrNotFound {
		d[key] = value
		return nil
	}
	return err
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	if err != nil {
		return err
	}
	delete(d, key)
	return nil
}
