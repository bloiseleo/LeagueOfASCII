package cache

import "fmt"

type KeyNotFound struct {
	Key string
}

func (knf *KeyNotFound) Error() string {
	return fmt.Sprintf("key %v was not found in cache", knf.Key)
}
