package cmp

type Sender interface {
	Send(v interface{}) error
}

type Receiver interface {
	Receiv() (interface{}, error)
}

type Communicator interface {
	Sender
	Receiver
}
