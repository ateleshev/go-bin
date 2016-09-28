package cmp

type Sender interface {
	Send(v interface{}) error
}

type Receiver interface {
	Receive() (interface{}, error)
}

type Communicator interface {
	Sender
	Receiver
}
