package service

import (
	"log"
	"net"
)

const BUFFER_SIZE = 1 << 16

type Agent struct {
	conn   net.Conn
	stoped bool
}

func NewAgent(conn net.Conn) *Agent {
	return &Agent{
		conn: conn,
	}
}

func (agent *Agent) Start() {
	go agent.read()
}

func (agent *Agent) read() {
	for !agent.stoped {
		buf := make([]byte, BUFFER_SIZE)

		n, err := agent.conn.Read(buf)
		if err != nil {
			log.Println("read error ", err)
		}

	}
}

func (agent *Agent) Send() {

}

func (agent *Agent) Stop() {
	agent.conn.Close()
	agent.stoped = true

}
