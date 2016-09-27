package data

type Getter interface {
	Get(key string, value interface{}) error
}

type Setter interface {
	Set(key string, value interface{}) error
}

type Client interface {
	Connect() error
	Close() error

	Info() string

	Getter
	Setter
}
