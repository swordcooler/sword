package client

import (
	"log"
	"net"
	"time"
)

type Client struct {
	addr string
	conn net.Conn
}

func NewClient(addr string) *Client {
	return &Client{
		addr: addr,
	}
}

func (c *Client) Start() {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		log.Println("connect error ", err)
	}

	c.conn = conn

	for {
		
		n, err := conn.Write([]byte("aabbccddeeffgg"))
		if err != nil {
			log.Println("write error ", err)
			continue
		}

		log.Println("write lenth ", n)

		time.Sleep(100)
	}
}
