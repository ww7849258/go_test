package maps

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	df, ok := d[key] // map 可返回两个值，第二个为bool类型，表示是否找到key
	if !ok {
		return "", ErrNotFound
	}

	return df, nil
}

func (d Dictionary) Add(k, v string) error {
	_, err := d.Search(k)

	switch err {
	case ErrNotFound:
		d[k] = v
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(k, v string) error {
	_, err := d.Search(k)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[k] = v
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(k string) {
	delete(d, k)
}
