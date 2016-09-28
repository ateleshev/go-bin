package cmp

type Sender interface {
	Send(interface{}) error
}

type Receiver interface {
	Receive() (interface{}, error)
}

type Communicator interface {
	Sender
	Receiver
}
