package client

import "testing"

func TestClient(t *testing.T) {
	for i := 0; i < 1; i++ {
		c := NewClient("192.168.1.189:8888")
		c.Start()
	}
}
