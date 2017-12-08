package client

import (
	"fmt"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	fmt.Println("start")
	for i := 0; i < 100; i++ {
		c := NewClient("192.168.1.189:8888")
		go c.Start()
	}

	time.Sleep(1000 * time.Second)
}

func TestPrintln(t *testing.T) {
	fmt.Println(1 << 24)
}
