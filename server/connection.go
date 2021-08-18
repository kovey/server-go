package server

import (
	"fmt"
	"net"
	"time"

	"github.com/kovey/logger-go/logger"
)

type Connection struct {
	fd          int64
	conn        net.Conn
	ip          string
	connectTime int64
	config      *Config
	packet      []byte
	curLen      int32
	packetLen   int32
	isClosed    bool
	header      *Header
}

func NewConnection(fd int64, conn net.Conn, c *Config) *Connection {
	return &Connection{
		fd: fd, conn: conn, ip: conn.RemoteAddr().String(), connectTime: time.Now().UnixNano() / 1e6, config: c, packet: make([]byte, 0, c.PackageMax),
		curLen: 0, packetLen: 0, isClosed: false,
	}
}

func (c *Connection) Fd() int64 {
	return c.fd
}

func (c *Connection) Read() ([]byte, *Header, error) {
	if c.packetLen > 0 && c.curLen >= c.packetLen {
		buf := c.packet[c.config.HeaderLen:c.packetLen]
		c.packet = c.packet[c.packetLen:]
		c.curLen -= c.packetLen
		header := c.header
		if c.curLen < c.config.HeaderLen {
			c.packetLen = 0
			return buf, header, nil
		}

		c.ParseHeader()
		return buf, header, nil
	}

	buf := make([]byte, 1024)
	n, err := c.conn.Read(buf)
	if err != nil {
		c.reset()
		return make([]byte, 0), nil, err
	}

	logger.Debug("receive len: %d\n", n)

	if n == 0 {
		c.reset()
		return make([]byte, 0), nil, fmt.Errorf("receive data is empty")
	}

	c.packet = append(c.packet, buf[:n]...)
	logger.Debug("current packet: %v, len: %d\n", c.packet, len(c.packet))
	c.curLen += int32(n)
	if c.packetLen == 0 {
		c.ParseHeader()
	}

	return make([]byte, 0), nil, nil
}

func (c *Connection) ParseHeader() {
	c.header = NewHeader()
	c.header.Parse(c.packet, c.config)
	c.packetLen = c.header.PackageLen
	logger.Debug("header decode, packetLen: %d, bodylen: %d\n", c.header.PackageLen, c.header.BodyLen)
}

func (c *Connection) reset() {
	c.curLen = 0
	c.packet = c.packet[:0]
	c.packetLen = 0
}

func (c *Connection) Write(buf []byte) (int, error) {
	return c.conn.Write(buf)
}

func (c *Connection) Close() error {
	if c.isClosed {
		return nil
	}

	c.isClosed = true
	return c.conn.Close()
}

func (c *Connection) Ip() string {
	return c.ip
}

func (c *Connection) ConnectTime() int64 {
	return c.connectTime
}

func (c *Connection) IsClosed() bool {
	return c.isClosed
}
