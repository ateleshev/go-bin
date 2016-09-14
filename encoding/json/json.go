package json

import "io"

const (
	hex = "0123456789abcdef"
)

var (
	EscapePrefix  = []byte(`\u00`)
	Utf8RuneError = []byte(`\ufffd`)
	Utf8SepPrefix = []byte(`\u202`)

	Null  = []byte(`null`)
	True  = []byte(`true`)
	False = []byte(`false`)
)

type Writer interface {
	io.Writer
	io.ByteWriter
}

type JsonWriterTo interface {
	JsonWriteTo(*JsonWriter) *JsonWriter
}
