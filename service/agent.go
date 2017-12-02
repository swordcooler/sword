package service

import "net"

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

	}
}

func (agent *Agent) Send() {

}

func (agent *Agent) Stop() {
	agent.stoped = true
}
