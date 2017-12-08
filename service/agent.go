package service

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"sync"
)

const BufferSize = 1 << 16
const HeaderSize = 1 << 3

var bufPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, BufferSize)
	},
}

// 每条链接一个agent
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
		buf := bufPool.Get().([]byte)
		n, err := agent.conn.Read(buf)
		log.Println(n)

		if err != nil {
			if err.Error() != "EOF" {
				log.Println("read error ", err)
			}

			agent.Stop()
			continue
		}

		if n <= 0 {
			continue
		}

	M:
		tag, length, err := Decode(buf[:HeaderSize])
		if err != nil {
			log.Println("decode header error ", err)
			continue
		}
		n -= HeaderSize
		buf = buf[HeaderSize:]
		if int(length) > n {
			continue
		}

		data := buf[:length]
		buf = buf[length:]
		n -= int(length)
		agent.handler(tag, data)

		if n > HeaderSize {
			log.Println("nianbao")
			goto M
		}

		bufPool.Put(buf)
	}
}

func (agent *Agent) Send() {

}

func (agent *Agent) Stop() {
	agent.conn.Close()
	agent.stoped = true

}

func (agent *Agent) handler(tag int32, data []byte) {
	log.Println("tag: ", tag, "data: ", string(data))
}

func Encode(tag int32, data string) ([]byte, error) {
	buf := new(bytes.Buffer)
	// 写入TAG
	if err := binary.Write(buf, binary.BigEndian, tag); err != nil {
		return nil, err
	}
	dataBuf := []byte(data)
	// 写入length
	if err := binary.Write(buf, binary.BigEndian, int32(len(dataBuf))); err != nil {
		return nil, err
	}
	// 写入数据
	if err := binary.Write(buf, binary.BigEndian, dataBuf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode(b []byte) (int32, int32, error) {
	buf := bytes.NewBuffer(b)
	var tag, length int32
	// 读取tag
	if err := binary.Read(buf, binary.BigEndian, &tag); err != nil {
		return 0, 0, err
	}
	// 读取length
	if err := binary.Read(buf, binary.BigEndian, &length); err != nil {
		return 0, 0, err
	}

	return tag, length, nil
}
