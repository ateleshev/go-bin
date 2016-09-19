package json

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

type JsonWriterTo interface {
	JsonWriteTo(*JsonWriter) *JsonWriter
}
