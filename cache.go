package cache

import (
	"errors"
)

type Cache struct {
	stor map[string]any
}

func New() *Cache {
	return &Cache{
		stor: make(map[string]any),
	}
}

func (c *Cache) Set(key string, value any) error {
	if err := validateKey(key); err != nil {
		return err
	}

	if err := validateValue(value); err != nil {
		return nil
	}

	c.stor[key] = value

	return nil
}

func (c Cache) Get(key string) (any, error) {
	if err := validateKey(key); err != nil {
		return nil, err
	}

	return c.stor[key], nil
}

func (c *Cache) Delete(key string) error {

	if err := validateKey(key); err != nil {
		return err
	}

	delete(c.stor, key)
	return nil
}

func validateKey(key string) error {
	if len(key) > 0 {
		return nil
	}
	return errors.New("keyIsEmptyErrText")
}

func validateValue(value any) error {
	if value != nil {
		return nil
	}
	return errors.New("valueIsEmptyErrText")
}
