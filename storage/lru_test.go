package storage

import (
	"fmt"
	"testing"
)

func TestNewLRUStore(t *testing.T) {
	store := NewLRUStore(64)
	store.Add("nous", "123456", "test")
	store.Add("nous", "12345", "test2")
	value, ok := store.Get("nous", "123456")
	if !ok {
		fmt.Println("Failed")
	}
	fmt.Println("value: ", value)

	l := store.Len("nous")
	fmt.Println("len: ", l)

	store.RemoveItem("nous", "123456")

	l = store.Len("nous")
	fmt.Println("len: ", l)

	store.Remove("nous")
	fmt.Println(len(store.Store))
}
