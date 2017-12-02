package client

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	fmt.Println("start")
	for i := 0; i < 1; i++ {
		c := NewClient("192.168.1.189:8888")
		c.Start()
	}
}

func TestPrintln(t *testing.T) {
	fmt.Println(1 << 24)
}
