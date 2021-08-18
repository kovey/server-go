package server

type EventInterface interface {
	Close(s *Server, fd int64)
	Connect(s *Server, fd int64)
	Receive(s *Server, e *Event)
}

type Event struct {
	header *Header
	body   []byte
	fd     int64
}

func NewEvent(header *Header, body []byte, fd int64) *Event {
	return &Event{header: header, body: body, fd: fd}
}

func (e *Event) Header() *Header {
	return e.header
}

func (e *Event) Body() []byte {
	return e.body
}

func (e *Event) Fd() int64 {
	return e.fd
}
