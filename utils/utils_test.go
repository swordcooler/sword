package utils

import (
	"fmt"
	"testing"
)

func TestNewUUID(t *testing.T) {
	fmt.Println("dddddddd")
	fmt.Println("ddddddddddd")

	for i := 0; i < 10; i++ {
		uuid := NewUUID()
		fmt.Println("new uuid ", uuid)
	}
}
