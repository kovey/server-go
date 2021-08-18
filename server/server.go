package server

import (
	"fmt"
	"net"

	"github.com/kovey/logger-go/logger"
)

type Server struct {
	host        string
	port        int32
	network     SocketType
	connections map[int64]*Connection
	maxFd       int64
	listener    net.Listener
	event       EventInterface
	config      *Config
}

type SocketType string

const (
	SOCKET_TCP SocketType = "tcp"
	SOCKET_UDP SocketType = "udp"
)

func NewServer(host string, port int32, network SocketType) *Server {
	return &Server{host: host, port: port, network: network, connections: make(map[int64]*Connection), maxFd: 0}
}

func (s *Server) Set(c *Config) {
	s.config = c
}

func (s *Server) SetEvent(event EventInterface) {
	s.event = event
}

func (s *Server) Start() {
	ln, err := net.Listen(string(s.network), fmt.Sprintf("%s:%d", s.host, s.port))
	logger.Debug("server listen on %s://%s:%d", s.network, s.host, s.port)
	if err != nil {
		panic(err)
	}

	s.listener = ln
	defer s.listener.Close()
	logger.GetInstance().Debug("server start")
	s.listen()
}

func (s *Server) listen() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			continue
		}

		s.maxFd++
		s.connections[s.maxFd] = NewConnection(s.maxFd, conn, s.config)
		go s.accept(s.maxFd, s.connections[s.maxFd])
	}
}

func (s *Server) accept(fd int64, conn *Connection) {
	defer func(s *Server, fd int64) {
		err := recover()
		if !logger.Panic(err) {
			return
		}

		s.Close(fd)
	}(s, fd)

	go func(s *Server, fd int64) {
		defer func(s *Server, fd int64) {
			err := recover()
			if !logger.Panic(err) {
				return
			}
			s.Close(fd)
		}(s, fd)

		s.event.Connect(s, fd)
	}(s, fd)

	defer conn.Close()
	for {
		if conn.IsClosed() {
			logger.GetInstance().Debug("client connections[%d] by closed", fd)
			break
		}

		buf, header, err := conn.Read()
		if err != nil {
			go func(s *Server, fd int64) {
				defer func() {
					err := recover()
					logger.Panic(err)
				}()
				s.Close(fd)
			}(s, fd)
			logger.Debug("close client connections[%d], receive error: %s", fd, err)
			break
		}

		if len(buf) == 0 {
			logger.Debug("receive data empty from client")
			continue
		}

		go func(s *Server, buf []byte, fd int64) {
			defer func(s *Server, fd int64) {
				err := recover()
				if !logger.Panic(err) {
					return
				}

				s.Close(fd)
			}(s, fd)
			s.event.Receive(s, NewEvent(header, buf, fd))
		}(s, buf, fd)
	}
}

func (s *Server) Close(fd int64) {
	conn, ok := s.connections[fd]
	if !ok {
		logger.GetInstance().Warning("connections[%d] is not exists", fd)
		return
	}

	delete(s.connections, fd)
	conn.Close()
	s.event.Close(s, fd)
}

func (s *Server) Send(buf []byte, fd int64) (int, error) {
	conn, ok := s.connections[fd]
	if !ok {
		return 0, fmt.Errorf("connections[%d] is not exists", fd)
	}

	return conn.Write(buf)

}
