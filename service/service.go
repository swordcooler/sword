package service

import (
	"log"
	"net"
	"sword/utils"
	"sync"
)

type Service struct {
	addr      string
	listener  net.Listener
	agents    map[string]*Agent
	stoped    bool
	stopping  bool
	agentLock sync.RWMutex
}

func NewService(addr string) *Service {
	return &Service{
		addr:   addr,
		agents: make(map[string]*Agent),
	}
}

func (s *Service) Start() {
	s.listen()
}

func (s *Service) listen() {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		panic(err)
	}

	s.listener = listener

	for !s.stoped && !s.stopping {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Println("accept error ", err)
			continue
		}

		uuid := utils.NewUUID()
		agent := NewAgent(conn)

		s.agentLock.Lock()
		s.agents[uuid] = agent
		s.agentLock.Unlock()

		agent.Start()

		//go s.handler(conn)
	}
}

func (s *Service) Stop() {
	s.agentLock.Lock()
	defer s.agentLock.Unlock()

	s.stopping = true
	agents := s.agents
	for _, agent := range agents {
		agent.Stop()
	}

	s.stoped = true
}

// func (s *Service) handler(conn net.Conn) {
// 	for !s.stoped && !s.stopping {
// 		buf := make([]byte, 64)

// 		n, err := conn.Read(buf)
// 		if err != nil {
// 			log.Println("read error ", err)
// 		}

// 		log.Println("length: ", n, " data: ", string(buf))
// 	}
// }
