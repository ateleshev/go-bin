package json

import "sync"
import "reflect"
import "strconv"
import "unicode/utf8"
import "github.com/ateleshev/go-bin/strings"

var jsonWriterPool = sync.Pool{ // {{{
	New: func() interface{} {
		return &JsonWriter{}
	},
} // }}}

func getJsonWriter() *JsonWriter { // {{{
	if instance := jsonWriterPool.Get(); instance != nil {
		return instance.(*JsonWriter)
	}

	return jsonWriterPool.New().(*JsonWriter)
} // }}}

func putJsonWriter(instance *JsonWriter) { // {{{
	instance.Reset()
	jsonWriterPool.Put(instance)
} // }}}

func NewJsonWriter(w Writer) *JsonWriter { // {{{
	jw := getJsonWriter()
	jw.writer = w

	return jw
} // }}}

type JsonWriter struct {
	htmlEscape bool
	writer     Writer
	errors     []error
}

func (this *JsonWriter) SetHtmlEscape(v bool) *JsonWriter { // {{{
	this.htmlEscape = v
	return this
} // }}}

func (this *JsonWriter) EnableHtmlEscape() *JsonWriter { // {{{
	this.htmlEscape = true
	return this
} // }}}

func (this *JsonWriter) Reset() *JsonWriter { // {{{
	this.htmlEscape = false
	this.writer = nil
	this.errors = this.errors[:0]
	return this
} // }}}

func (this *JsonWriter) Release() *JsonWriter { // {{{
	putJsonWriter(this)
	return this
} // }}}

func (this *JsonWriter) Errors() []error { // {{{
	return this.errors
} // }}}

func (this *JsonWriter) HasErrors() bool { // {{{
	return len(this.errors) > 0
} // }}}

func (this *JsonWriter) LastError() error { // {{{
	if len(this.errors) > 0 {
		return this.errors[len(this.errors)-1]
	}
	return nil
} // }}}

func (this *JsonWriter) errorHandler(n int, err error) *JsonWriter { // {{{
	if err != nil {
		this.errors = append(this.errors, err)
	}
	return this
} // }}}

func (this *JsonWriter) writeNull() *JsonWriter { // {{{
	return this.write(Null)
} // }}}

func (this *JsonWriter) writeBool(v bool) *JsonWriter { // {{{
	if v {
		return this.write(True)
	}

	return this.write(False)
} // }}}

func (this *JsonWriter) writeInt64(v int64) *JsonWriter { // {{{
	buf := newNumericValueBuffer()
	defer buf.Release()
	buf.B = strconv.AppendInt(buf.B, v, 10)
	return this.write(buf.B[0:len(buf.B)])
} // }}}

func (this *JsonWriter) writeUint64(v uint64) *JsonWriter { // {{{
	buf := newNumericValueBuffer()
	defer buf.Release()
	buf.B = strconv.AppendUint(buf.B, v, 10)
	return this.write(buf.B[0:len(buf.B)])
} // }}}

func (this *JsonWriter) writeFloat64(v float64, bitSize int) *JsonWriter { // {{{
	buf := newNumericValueBuffer()
	defer buf.Release()
	buf.B = strconv.AppendFloat(buf.B, v, 'g', -1, bitSize)
	return this.write(buf.B[0:len(buf.B)])
} // }}}

func (this *JsonWriter) writeByte(b byte) *JsonWriter { // {{{
	return this.errorHandler(1, this.writer.WriteByte(b))
} // }}}

func (this *JsonWriter) write(v []byte) *JsonWriter { // {{{
	return this.errorHandler(this.writer.Write(v))
} // }}}

/**
 * Details: https://golang.org/src/encoding/json/encode.go
 *   func (e *encodeState) stringBytes(s []byte, escapeHTML bool) int
 */
func (this *JsonWriter) escape(v []byte) *JsonWriter { // {{{
	// Start, Current index
	var start, i int
	for i < len(v) {
		if v[i] < utf8.RuneSelf {
			switch v[i] {
			case '\\', '"':
				this.write(v[start:i]).writeByte('\\').writeByte(v[i])
				goto encode_continue
			case '\t':
				this.write(v[start:i]).writeByte('\\').writeByte('t')
				goto encode_continue
			case '\r':
				this.write(v[start:i]).writeByte('\\').writeByte('r')
				goto encode_continue
			case '\n':
				this.write(v[start:i]).writeByte('\\').writeByte('n')
				goto encode_continue
			case '<', '>', '&':
				if this.htmlEscape {
					goto encode_bytes
				}
				break
			default:
				// This encodes bytes < 0x20 except for \t, \n and \r.
				// If escapeHTML is set, it also escapes <, >, and &
				// because they can lead to security holes when
				// user-controlled strings are rendered into JSON
				// and served to some browsers.
				if v[i] < 0x20 {
					goto encode_bytes
				}
				break
			}

			i++
			continue

		encode_bytes:
			this.write(v[start:i]).write(EscapePrefix)
			this.writeByte(hex[v[i]>>4]).writeByte(hex[v[i]&0xF])
		encode_continue:
			i++
			start = i
			continue
		}

		r, size := utf8.DecodeRune(v[i:])

		if r == utf8.RuneError && size == 1 {
			this.write(v[start:i]).write(Utf8RuneError)
			i += size
			start = i
			continue
		}

		// U+2028 is LINE SEPARATOR.
		// U+2029 is PARAGRAPH SEPARATOR.
		// They are both technically valid characters in JSON strings,
		// but don't work in JSONP, which has to be evaluated as JavaScript,
		// and can lead to security holes there. It is valid JSON to
		// escape them, so we do so unconditionally.
		// See http://timelessrepo.com/json-isnt-a-javascript-subset for discussion.
		if r == '\u2028' || r == '\u2029' {
			this.write(v[start:i]).write(Utf8SepPrefix)
			this.writeByte(hex[r&0xF])
			i += size
			start = i
			continue
		}

		i += size
	}

	return this.write(v[start:i])
} // }}}

func (this *JsonWriter) ObjOpen() *JsonWriter { // {{{
	return this.writeByte('{')
} // }}}

func (this *JsonWriter) ObjClose() *JsonWriter { // {{{
	return this.writeByte('}')
} // }}}

func (this *JsonWriter) ArrOpen() *JsonWriter { // {{{
	return this.writeByte('[')
} // }}}

func (this *JsonWriter) ArrClose() *JsonWriter { // {{{
	return this.writeByte(']')
} // }}}

func (this *JsonWriter) Next() *JsonWriter { // {{{
	return this.writeByte(',')
} // }}}

func (this *JsonWriter) NextIf(v bool) *JsonWriter { // {{{
	if v {
		return this.writeByte(',')
	}
	return this
} // }}}

func (this *JsonWriter) Sep() *JsonWriter { // {{{
	return this.writeByte(':')
} // }}}

/**
 * Add value in double quotes:
 *   some value -> "some value"
 */
func (this *JsonWriter) Quoted(v []byte) *JsonWriter { // {{{
	return this.writeByte('"').escape(v).writeByte('"')
} // }}}

/**
 * Add value without double quotes:
 *   12345 -> 12345  (int)
 *   12.45 -> 12.45  (float64)
 *   false -> false  (bool)
 *   true  -> true   (bool)
 *   null  -> null   (nil)
 */
func (this *JsonWriter) Unquoted(v []byte) *JsonWriter { // {{{
	return this.write(v)
} // }}}

func (this *JsonWriter) NullValue() *JsonWriter { // {{{
	return this.writeNull()
} // }}}

func (this *JsonWriter) BoolValue(v bool) *JsonWriter { // {{{
	return this.writeBool(v)
} // }}}

func (this *JsonWriter) BoolPtrValue(v *bool) *JsonWriter { // {{{
	return this.BoolValue(*v)
} // }}}

func (this *JsonWriter) BytesValue(v []byte) *JsonWriter { // {{{
	return this.Quoted(v)
} // }}}

func (this *JsonWriter) StringValue(v string) *JsonWriter { // {{{
	return this.Quoted(strings.ToBytes(v))
} // }}}

func (this *JsonWriter) StringPtrValue(v *string) *JsonWriter { // {{{
	return this.Quoted(strings.PtrToBytes(v))
} // }}}

func (this *JsonWriter) Float32Value(v float32) *JsonWriter { // {{{
	return this.writeFloat64(float64(v), 32)
} // }}}

func (this *JsonWriter) Float32PtrValue(v *float32) *JsonWriter { // {{{
	return this.Float32Value(*v)
} // }}}

func (this *JsonWriter) Float64Value(v float64) *JsonWriter { // {{{
	return this.writeFloat64(v, 64)
} // }}}

func (this *JsonWriter) Float64PtrValue(v *float64) *JsonWriter { // {{{
	return this.Float64Value(*v)
} // }}}

func (this *JsonWriter) IntValue(v int) *JsonWriter { // {{{
	return this.writeInt64(int64(v))
} // }}}

func (this *JsonWriter) IntPtrValue(v *int) *JsonWriter { // {{{
	return this.IntValue(*v)
} // }}}

func (this *JsonWriter) Int8Value(v int8) *JsonWriter { // {{{
	return this.writeInt64(int64(v))
} // }}}

func (this *JsonWriter) Int8PtrValue(v *int8) *JsonWriter { // {{{
	return this.Int8Value(*v)
} // }}}

func (this *JsonWriter) Int16Value(v int16) *JsonWriter { // {{{
	return this.writeInt64(int64(v))
} // }}}

func (this *JsonWriter) Int16PtrValue(v *int16) *JsonWriter { // {{{
	return this.Int16Value(*v)
} // }}}

func (this *JsonWriter) Int32Value(v int32) *JsonWriter { // {{{
	return this.writeInt64(int64(v))
} // }}}

func (this *JsonWriter) Int32PtrValue(v *int32) *JsonWriter { // {{{
	return this.Int32Value(*v)
} // }}}

func (this *JsonWriter) Int64Value(v int64) *JsonWriter { // {{{
	return this.writeInt64(v)
} // }}}

func (this *JsonWriter) Int64PtrValue(v *int64) *JsonWriter { // {{{
	return this.Int64Value(*v)
} // }}}

func (this *JsonWriter) UintValue(v uint) *JsonWriter { // {{{
	return this.writeUint64(uint64(v))
} // }}}

func (this *JsonWriter) UintPtrValue(v *uint) *JsonWriter { // {{{
	return this.UintValue(*v)
} // }}}

func (this *JsonWriter) Uint8Value(v uint8) *JsonWriter { // {{{
	return this.writeUint64(uint64(v))
} // }}}

func (this *JsonWriter) Uint8PtrValue(v *uint8) *JsonWriter { // {{{
	return this.Uint8Value(*v)
} // }}}

func (this *JsonWriter) Uint16Value(v uint16) *JsonWriter { // {{{
	return this.writeUint64(uint64(v))
} // }}}

func (this *JsonWriter) Uint16PtrValue(v *uint16) *JsonWriter { // {{{
	return this.Uint16Value(*v)
} // }}}

func (this *JsonWriter) Uint32Value(v uint32) *JsonWriter { // {{{
	return this.writeUint64(uint64(v))
} // }}}

func (this *JsonWriter) Uint32PtrValue(v *uint32) *JsonWriter { // {{{
	return this.Uint32Value(*v)
} // }}}

func (this *JsonWriter) Uint64Value(v uint64) *JsonWriter { // {{{
	return this.writeUint64(v)
} // }}}

func (this *JsonWriter) Uint64PtrValue(v *uint64) *JsonWriter { // {{{
	return this.Uint64Value(*v)
} // }}}

func (this *JsonWriter) Value(v interface{}) *JsonWriter { // {{{
	if v == nil || reflect.ValueOf(v).IsNil() {
		return this.NullValue()
	}

	if instance, ok := v.(JsonWriterTo); ok {
		return instance.JsonWriteTo(this)
	}

	switch v.(type) {
	case bool:
		return this.BoolValue(v.(bool))
	case *bool:
		return this.BoolPtrValue(v.(*bool))
	case []byte:
		return this.BytesValue(v.([]byte))
	case string:
		return this.StringValue(v.(string))
	case *string:
		return this.StringPtrValue(v.(*string))
	case float32:
		return this.Float32Value(v.(float32))
	case *float32:
		return this.Float32PtrValue(v.(*float32))
	case float64:
		return this.Float64Value(v.(float64))
	case *float64:
		return this.Float64PtrValue(v.(*float64))
	case int:
		return this.IntValue(v.(int))
	case *int:
		return this.IntPtrValue(v.(*int))
	case int8:
		return this.Int8Value(v.(int8))
	case *int8:
		return this.Int8PtrValue(v.(*int8))
	case int16:
		return this.Int16Value(v.(int16))
	case *int16:
		return this.Int16PtrValue(v.(*int16))
	case int32:
		return this.Int32Value(v.(int32))
	case *int32:
		return this.Int32PtrValue(v.(*int32))
	case int64:
		return this.Int64Value(v.(int64))
	case *int64:
		return this.Int64PtrValue(v.(*int64))
	case uint:
		return this.UintValue(v.(uint))
	case *uint:
		return this.UintPtrValue(v.(*uint))
	case uint8:
		return this.Uint8Value(v.(uint8))
	case *uint8:
		return this.Uint8PtrValue(v.(*uint8))
	case uint16:
		return this.Uint16Value(v.(uint16))
	case *uint16:
		return this.Uint16PtrValue(v.(*uint16))
	case uint32:
		return this.Uint32Value(v.(uint32))
	case *uint32:
		return this.Uint32PtrValue(v.(*uint32))
	case uint64:
		return this.Uint64Value(v.(uint64))
	case *uint64:
		return this.Uint64PtrValue(v.(*uint64))
	default:
		panic("json: unsupported type:" + reflect.TypeOf(v).String())
	}

	return this
} // }}}

/**
 * Add element of object
 *
 * n - element name
 * v - element value
 * f - is first element
 */
func (this *JsonWriter) ObjElm(n, v interface{}, f bool) *JsonWriter { // {{{
	return this.ObjElmName(n, f).ObjElmVal(v)
} // }}}

func (this *JsonWriter) ObjElmName(n interface{}, f bool) *JsonWriter { // {{{
	if !f {
		this.Next()
	}
	return this.Value(n).Sep()
} // }}}

func (this *JsonWriter) ObjElmVal(v interface{}) *JsonWriter { // {{{
	return this.Value(v)
} // }}}

/**
 * Add element of array
 *
 * v - element value
 * f - is first element
 */
func (this *JsonWriter) ArrElm(v interface{}, f bool) *JsonWriter { // {{{
	if !f {
		this.Next()
	}
	return this.Value(v)
} // }}}
