package json_test

import "testing"
import "bytes"
import "strconv"
import "github.com/ateleshev/go-bin/encoding/json"

/**
 * ==[ Tests ]==
 *
 * go test -v -run=JsonWriter_
 */

// --[ null ]--

func TestJsonWriter_Null(t *testing.T) { // {{{
	bf := bytes.NewBuffer(make([]byte, 0, 4))
	jw := json.NewJsonWriter(bf)

	jw.NullValue()
	if bf.String() != "null" {
		t.Fatal("Incorrect result for: null !=", bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

// --[ /null ]--

// --[ bool ]--

func TestJsonWriter_Bool(t *testing.T) { // {{{
	var v bool
	bf := bytes.NewBuffer(make([]byte, 0, 5))
	jw := json.NewJsonWriter(bf)

	jw.BoolValue(v)
	if bf.String() != "false" {
		t.Fatal("Incorrect result for: false !=", bf.String())
	}
	bf.Reset()

	v = true
	jw.BoolValue(v)
	if bf.String() != "true" {
		t.Fatal("Incorrect result for: true !=", bf.String())
	}
	bf.Reset()

	v = false
	jw.BoolValue(v)
	if bf.String() != "false" {
		t.Fatal("Incorrect result for: false !=", bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

func TestJsonWriter_BoolPtr(t *testing.T) { // {{{
	var v bool
	bf := bytes.NewBuffer(make([]byte, 0, 5))
	jw := json.NewJsonWriter(bf)

	jw.BoolPtrValue(&v)
	if bf.String() != "false" {
		t.Fatal("Incorrect result for: false !=", bf.String())
	}
	bf.Reset()

	v = true
	jw.BoolPtrValue(&v)
	if bf.String() != "true" {
		t.Fatal("Incorrect result for: true !=", bf.String())
	}
	bf.Reset()

	v = false
	jw.BoolPtrValue(&v)
	if bf.String() != "false" {
		t.Fatal("Incorrect result for: false !=", bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

// --[ /bool ]--

// --[ string ]--

func TestJsonWriter_String(t *testing.T) { // {{{
	var v string
	bf := bytes.NewBuffer(make([]byte, 0, 1024))
	jw := json.NewJsonWriter(bf)

	// English alphabet
	v = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	jw.StringValue(v)
	if bf.String() != `"`+v+`"` {
		t.Fatal(`Incorrect result for:"`, v, `"!=`, bf.String())
	}
	bf.Reset()

	// Ukrainian alphabet
	v = "АаБбВвГгҐґДдЕеЄєЖжЗзИиІіЇїЙйКкЛлМмНнОоПпРрСсТтУуФфХхЦцЧчШшЩщЬьЮюЯя"
	jw.StringValue(v)
	if bf.String() != `"`+v+`"` {
		t.Fatal(`Incorrect result for:"`, v, `"!=`, bf.String())
	}
	bf.Reset()

	v = "0123456789"
	jw.StringValue(v)
	if bf.String() != `"`+v+`"` {
		t.Fatal(`Incorrect result for:"`, v, `"!=`, bf.String())
	}
	bf.Reset()

	v = "\\"
	jw.StringValue(v)
	if bf.String() != `"\\"` {
		t.Fatal(`Incorrect result for: "\\" !=`, bf.String())
	}
	bf.Reset()

	v = "\""
	jw.StringValue(v)
	if bf.String() != `"\""` {
		t.Fatal(`Incorrect result for: "\"" !=`, bf.String())
	}
	bf.Reset()

	v = "\t"
	jw.StringValue(v)
	if bf.String() != `"\t"` {
		t.Fatal(`Incorrect result for: "\t" !=`, bf.String())
	}
	bf.Reset()

	v = "\r"
	jw.StringValue(v)
	if bf.String() != `"\r"` {
		t.Fatal(`Incorrect result for: "\r" !=`, bf.String())
	}
	bf.Reset()

	v = "\n"
	jw.StringValue(v)
	if bf.String() != `"\n"` {
		t.Fatal(`Incorrect result for: "\n" !=`, bf.String())
	}
	bf.Reset()

	v = "\b"
	jw.StringValue(v)
	if bf.String() != `"\u0008"` {
		t.Fatal(`Incorrect result for: "\u0008" !=`, bf.String())
	}
	bf.Reset()

	v = "\f"
	jw.StringValue(v)
	if bf.String() != `"\u000c"` {
		t.Fatal(`Incorrect result for: "\u000c" !=`, bf.String())
	}
	bf.Reset()

	v = "\u2028"
	jw.StringValue(v)
	if bf.String() != `"\u2028"` {
		t.Fatal(`Incorrect result for: "\u2028" !=`, bf.String())
	}
	bf.Reset()

	v = "\u2029"
	jw.StringValue(v)
	if bf.String() != `"\u2029"` {
		t.Fatal(`Incorrect result for: "\u2029" !=`, bf.String())
	}
	bf.Reset()

	v = "<"
	jw.StringValue(v)
	if bf.String() != `"<"` {
		t.Fatal(`Incorrect result for: "<" !=`, bf.String())
	}
	bf.Reset()

	v = ">"
	jw.StringValue(v)
	if bf.String() != `">"` {
		t.Fatal(`Incorrect result for: ">" !=`, bf.String())
	}
	bf.Reset()

	v = "&"
	jw.StringValue(v)
	if bf.String() != `"&"` {
		t.Fatal(`Incorrect result for: "&" !=`, bf.String())
	}
	bf.Reset()

	// -- HTML escape --

	jw.EnableHtmlEscape()

	v = "<"
	jw.StringValue(v)
	if bf.String() != `"\u003c"` {
		t.Fatal(`Incorrect result for: "<" !=`, bf.String())
	}
	bf.Reset()

	v = ">"
	jw.StringValue(v)
	if bf.String() != `"\u003e"` {
		t.Fatal(`Incorrect result for: ">" !=`, bf.String())
	}
	bf.Reset()

	v = "&"
	jw.StringValue(v)
	if bf.String() != `"\u0026"` {
		t.Fatal(`Incorrect result for: "\u0026" !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

func TestJsonWriter_StringPtr(t *testing.T) { // {{{
	var v string
	bf := bytes.NewBuffer(make([]byte, 0, 1024))
	jw := json.NewJsonWriter(bf)

	// English alphabet
	v = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	jw.StringPtrValue(&v)
	if bf.String() != `"`+v+`"` {
		t.Fatal(`Incorrect result for:"`, v, `"!=`, bf.String())
	}
	bf.Reset()

	// Ukrainian alphabet
	v = "АаБбВвГгҐґДдЕеЄєЖжЗзИиІіЇїЙйКкЛлМмНнОоПпРрСсТтУуФфХхЦцЧчШшЩщЬьЮюЯя"
	jw.StringPtrValue(&v)
	if bf.String() != `"`+v+`"` {
		t.Fatal(`Incorrect result for:"`, v, `"!=`, bf.String())
	}
	bf.Reset()

	v = "0123456789"
	jw.StringPtrValue(&v)
	if bf.String() != `"`+v+`"` {
		t.Fatal(`Incorrect result for:"`, v, `"!=`, bf.String())
	}
	bf.Reset()

	v = "\\"
	jw.StringPtrValue(&v)
	if bf.String() != `"\\"` {
		t.Fatal(`Incorrect result for: "\\" !=`, bf.String())
	}
	bf.Reset()

	v = "\""
	jw.StringPtrValue(&v)
	if bf.String() != `"\""` {
		t.Fatal(`Incorrect result for: "\"" !=`, bf.String())
	}
	bf.Reset()

	v = "\t"
	jw.StringPtrValue(&v)
	if bf.String() != `"\t"` {
		t.Fatal(`Incorrect result for: "\t" !=`, bf.String())
	}
	bf.Reset()

	v = "\r"
	jw.StringPtrValue(&v)
	if bf.String() != `"\r"` {
		t.Fatal(`Incorrect result for: "\r" !=`, bf.String())
	}
	bf.Reset()

	v = "\n"
	jw.StringPtrValue(&v)
	if bf.String() != `"\n"` {
		t.Fatal(`Incorrect result for: "\n" !=`, bf.String())
	}
	bf.Reset()

	v = "\b"
	jw.StringPtrValue(&v)
	if bf.String() != `"\u0008"` {
		t.Fatal(`Incorrect result for: "\u0008" !=`, bf.String())
	}
	bf.Reset()

	v = "\f"
	jw.StringPtrValue(&v)
	if bf.String() != `"\u000c"` {
		t.Fatal(`Incorrect result for: "\u000c" !=`, bf.String())
	}
	bf.Reset()

	v = "\u2028"
	jw.StringPtrValue(&v)
	if bf.String() != `"\u2028"` {
		t.Fatal(`Incorrect result for: "\u2028" !=`, bf.String())
	}
	bf.Reset()

	v = "\u2029"
	jw.StringPtrValue(&v)
	if bf.String() != `"\u2029"` {
		t.Fatal(`Incorrect result for: "\u2029" !=`, bf.String())
	}
	bf.Reset()

	v = "<"
	jw.StringPtrValue(&v)
	if bf.String() != `"<"` {
		t.Fatal(`Incorrect result for: "<" !=`, bf.String())
	}
	bf.Reset()

	v = ">"
	jw.StringPtrValue(&v)
	if bf.String() != `">"` {
		t.Fatal(`Incorrect result for: ">" !=`, bf.String())
	}
	bf.Reset()

	v = "&"
	jw.StringPtrValue(&v)
	if bf.String() != `"&"` {
		t.Fatal(`Incorrect result for: "&" !=`, bf.String())
	}
	bf.Reset()

	// -- HTML escape --

	jw.EnableHtmlEscape()

	v = "<"
	jw.StringPtrValue(&v)
	if bf.String() != `"\u003c"` {
		t.Fatal(`Incorrect result for: "<" !=`, bf.String())
	}
	bf.Reset()

	v = ">"
	jw.StringPtrValue(&v)
	if bf.String() != `"\u003e"` {
		t.Fatal(`Incorrect result for: ">" !=`, bf.String())
	}
	bf.Reset()

	v = "&"
	jw.StringPtrValue(&v)
	if bf.String() != `"\u0026"` {
		t.Fatal(`Incorrect result for: "\u0026" !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

// --[ /string ]--

// --[ int ]--

/**
 * int8 is the set of all signed 8-bit integers. Range: -128 through 127.
 */

func TestJsonWriter_Int8(t *testing.T) { // {{{
	var v int8
	bf := bytes.NewBuffer(make([]byte, 0, 4))
	jw := json.NewJsonWriter(bf)

	jw.Int8Value(v)
	if bf.String() != `0` {
		t.Fatal(`Incorrect result for: 0 !=`, bf.String())
	}
	bf.Reset()

	v = 127
	jw.Int8Value(v)
	if bf.String() != `127` {
		t.Fatal(`Incorrect result for: 127 !=`, bf.String())
	}
	bf.Reset()

	v = -128
	jw.Int8Value(v)
	if bf.String() != `-128` {
		t.Fatal(`Incorrect result for: -128 !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

func TestJsonWriter_Int8Ptr(t *testing.T) { // {{{
	var v int8
	bf := bytes.NewBuffer(make([]byte, 0, 4))
	jw := json.NewJsonWriter(bf)

	jw.Int8PtrValue(&v)
	if bf.String() != `0` {
		t.Fatal(`Incorrect result for: 0 !=`, bf.String())
	}
	bf.Reset()

	v = 127
	jw.Int8PtrValue(&v)
	if bf.String() != `127` {
		t.Fatal(`Incorrect result for: 127 !=`, bf.String())
	}
	bf.Reset()

	v = -128
	jw.Int8PtrValue(&v)
	if bf.String() != `-128` {
		t.Fatal(`Incorrect result for: -128 !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

/**
 * int16 is the set of all signed 16-bit integers. Range: -32768 through 32767.
 */

func TestJsonWriter_Int16(t *testing.T) { // {{{
	var v int16
	bf := bytes.NewBuffer(make([]byte, 0, 6))
	jw := json.NewJsonWriter(bf)

	jw.Int16Value(v)
	if bf.String() != `0` {
		t.Fatal(`Incorrect result for: 0 !=`, bf.String())
	}
	bf.Reset()

	v = 32767
	jw.Int16Value(v)
	if bf.String() != `32767` {
		t.Fatal(`Incorrect result for: 32767 !=`, bf.String())
	}
	bf.Reset()

	v = -32768
	jw.Int16Value(v)
	if bf.String() != `-32768` {
		t.Fatal(`Incorrect result for: -32768 !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

func TestJsonWriter_Int16Ptr(t *testing.T) { // {{{
	var v int16
	bf := bytes.NewBuffer(make([]byte, 0, 6))
	jw := json.NewJsonWriter(bf)

	jw.Int16PtrValue(&v)
	if bf.String() != `0` {
		t.Fatal(`Incorrect result for: 0 !=`, bf.String())
	}
	bf.Reset()

	v = 32767
	jw.Int16PtrValue(&v)
	if bf.String() != `32767` {
		t.Fatal(`Incorrect result for: 32767 !=`, bf.String())
	}
	bf.Reset()

	v = -32768
	jw.Int16PtrValue(&v)
	if bf.String() != `-32768` {
		t.Fatal(`Incorrect result for: -32768 !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

/**
 * int32 is the set of all signed 32-bit integers. Range: -2147483648 through 2147483647.
 */

func TestJsonWriter_Int32(t *testing.T) { // {{{
	var v int32
	bf := bytes.NewBuffer(make([]byte, 0, 11))
	jw := json.NewJsonWriter(bf)

	jw.Int32Value(v)
	if bf.String() != `0` {
		t.Fatal(`Incorrect result for: 0 !=`, bf.String())
	}
	bf.Reset()

	v = 2147483647
	jw.Int32Value(v)
	if bf.String() != `2147483647` {
		t.Fatal(`Incorrect result for: 2147483647 !=`, bf.String())
	}
	bf.Reset()

	v = -2147483648
	jw.Int32Value(v)
	if bf.String() != `-2147483648` {
		t.Fatal(`Incorrect result for: -2147483648 !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

func TestJsonWriter_Int32Ptr(t *testing.T) { // {{{
	var v int32
	bf := bytes.NewBuffer(make([]byte, 0, 11))
	jw := json.NewJsonWriter(bf)

	jw.Int32PtrValue(&v)
	if bf.String() != `0` {
		t.Fatal(`Incorrect result for: 0 !=`, bf.String())
	}
	bf.Reset()

	v = 2147483647
	jw.Int32PtrValue(&v)
	if bf.String() != `2147483647` {
		t.Fatal(`Incorrect result for: 2147483647 !=`, bf.String())
	}
	bf.Reset()

	v = -2147483648
	jw.Int32PtrValue(&v)
	if bf.String() != `-2147483648` {
		t.Fatal(`Incorrect result for: -2147483648 !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

/**
 * int64 is the set of all signed 64-bit integers. Range: -9223372036854775808 through 9223372036854775807.
 */

func TestJsonWriter_Int64(t *testing.T) { // {{{
	if strconv.IntSize != 64 {
		t.Skip(`Your system is not support 64-bit`)
	}

	var v int64
	bf := bytes.NewBuffer(make([]byte, 0, 20))
	jw := json.NewJsonWriter(bf)

	jw.Int64Value(v)
	if bf.String() != `0` {
		t.Fatal(`Incorrect result for: 0 !=`, bf.String())
	}
	bf.Reset()

	v = 9223372036854775807
	jw.Int64Value(v)
	if bf.String() != `9223372036854775807` {
		t.Fatal(`Incorrect result for: 9223372036854775807 !=`, bf.String())
	}
	bf.Reset()

	v = -9223372036854775808
	jw.Int64Value(v)
	if bf.String() != `-9223372036854775808` {
		t.Fatal(`Incorrect result for: -9223372036854775808 !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

func TestJsonWriter_Int64Ptr(t *testing.T) { // {{{
	if strconv.IntSize != 64 {
		t.Skip(`Your system is not support 64-bit`)
	}

	var v int64
	bf := bytes.NewBuffer(make([]byte, 0, 20))
	jw := json.NewJsonWriter(bf)

	jw.Int64PtrValue(&v)
	if bf.String() != `0` {
		t.Fatal(`Incorrect result for: 0 !=`, bf.String())
	}
	bf.Reset()

	v = 9223372036854775807
	jw.Int64PtrValue(&v)
	if bf.String() != `9223372036854775807` {
		t.Fatal(`Incorrect result for: 9223372036854775807 !=`, bf.String())
	}
	bf.Reset()

	v = -9223372036854775808
	jw.Int64PtrValue(&v)
	if bf.String() != `-9223372036854775808` {
		t.Fatal(`Incorrect result for: -9223372036854775808 !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

/**
 * int is a signed integer type that is at least 32 bits in size. It is a distinct type, however, and not an alias for, say, int32.
 */

func TestJsonWriter_Int(t *testing.T) { // {{{
	var v int
	if strconv.IntSize < 64 { // 32-bit
		bf := bytes.NewBuffer(make([]byte, 0, 11))
		jw := json.NewJsonWriter(bf)

		jw.IntValue(v)
		if bf.String() != `0` {
			t.Fatal(`Incorrect result for: 0 !=`, bf.String())
		}
		bf.Reset()

		v = 2147483647
		jw.IntValue(v)
		if bf.String() != `2147483647` {
			t.Fatal(`Incorrect result for: 2147483647 !=`, bf.String())
		}
		bf.Reset()

		v = -2147483648
		jw.IntValue(v)
		if bf.String() != `-2147483648` {
			t.Fatal(`Incorrect result for: -2147483648 !=`, bf.String())
		}
		bf.Reset()

		jw.Release()
	} else { // 64-bit
		bf := bytes.NewBuffer(make([]byte, 0, 20))
		jw := json.NewJsonWriter(bf)

		jw.IntValue(v)
		if bf.String() != `0` {
			t.Fatal(`Incorrect result for: 0 !=`, bf.String())
		}
		bf.Reset()

		v = 9223372036854775807
		jw.IntValue(v)
		if bf.String() != `9223372036854775807` {
			t.Fatal(`Incorrect result for: 9223372036854775807 !=`, bf.String())
		}
		bf.Reset()

		v = -9223372036854775808
		jw.IntValue(v)
		if bf.String() != `-9223372036854775808` {
			t.Fatal(`Incorrect result for: -9223372036854775808 !=`, bf.String())
		}
		bf.Reset()

		jw.Release()
	}
} // }}}

func TestJsonWriter_IntPtr(t *testing.T) { // {{{
	var v int
	if strconv.IntSize != 64 {
		bf := bytes.NewBuffer(make([]byte, 0, 11))
		jw := json.NewJsonWriter(bf)

		jw.IntPtrValue(&v)
		if bf.String() != `0` {
			t.Fatal(`Incorrect result for: 0 !=`, bf.String())
		}
		bf.Reset()

		v = 2147483647
		jw.IntPtrValue(&v)
		if bf.String() != `2147483647` {
			t.Fatal(`Incorrect result for: 2147483647 !=`, bf.String())
		}
		bf.Reset()

		v = -2147483648
		jw.IntPtrValue(&v)
		if bf.String() != `-2147483648` {
			t.Fatal(`Incorrect result for: -2147483648 !=`, bf.String())
		}
		bf.Reset()

		jw.Release()
	} else {
		bf := bytes.NewBuffer(make([]byte, 0, 20))
		jw := json.NewJsonWriter(bf)

		jw.IntPtrValue(&v)
		if bf.String() != `0` {
			t.Fatal(`Incorrect result for: 0 !=`, bf.String())
		}
		bf.Reset()

		v = 9223372036854775807
		jw.IntPtrValue(&v)
		if bf.String() != `9223372036854775807` {
			t.Fatal(`Incorrect result for: 9223372036854775807 !=`, bf.String())
		}
		bf.Reset()

		v = -9223372036854775808
		jw.IntPtrValue(&v)
		if bf.String() != `-9223372036854775808` {
			t.Fatal(`Incorrect result for: -9223372036854775808 !=`, bf.String())
		}
		bf.Reset()

		jw.Release()
	}
} // }}}

// --[ /int ]--

// --[ uint ]--

/**
 * uint8 is the set of all unsigned 8-bit integers. Range: 0 through 255.
 */

func TestJsonWriter_Uint8(t *testing.T) { // {{{
	var v uint8
	bf := bytes.NewBuffer(make([]byte, 0, 3))
	jw := json.NewJsonWriter(bf)

	jw.Uint8Value(v)
	if bf.String() != `0` {
		t.Fatal(`Incorrect result for: 0 !=`, bf.String())
	}
	bf.Reset()

	v = 255
	jw.Uint8Value(v)
	if bf.String() != `255` {
		t.Fatal(`Incorrect result for: 255 !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

func TestJsonWriter_Uint8Ptr(t *testing.T) { // {{{
	var v uint8
	bf := bytes.NewBuffer(make([]byte, 0, 3))
	jw := json.NewJsonWriter(bf)

	jw.Uint8PtrValue(&v)
	if bf.String() != `0` {
		t.Fatal(`Incorrect result for: 0 !=`, bf.String())
	}
	bf.Reset()

	v = 255
	jw.Uint8PtrValue(&v)
	if bf.String() != `255` {
		t.Fatal(`Incorrect result for: 255 !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

/**
 * uint16 is the set of all unsigned 16-bit integers. Range: 0 through 65535.
 */

func TestJsonWriter_Uint16(t *testing.T) { // {{{
	var v uint16
	bf := bytes.NewBuffer(make([]byte, 0, 5))
	jw := json.NewJsonWriter(bf)

	jw.Uint16Value(v)
	if bf.String() != `0` {
		t.Fatal(`Incorrect result for: 0 !=`, bf.String())
	}
	bf.Reset()

	v = 65535
	jw.Uint16Value(v)
	if bf.String() != `65535` {
		t.Fatal(`Incorrect result for: 65535 !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

func TestJsonWriter_Uint16Ptr(t *testing.T) { // {{{
	var v uint16
	bf := bytes.NewBuffer(make([]byte, 0, 5))
	jw := json.NewJsonWriter(bf)

	jw.Uint16PtrValue(&v)
	if bf.String() != `0` {
		t.Fatal(`Incorrect result for: 0 !=`, bf.String())
	}
	bf.Reset()

	v = 65535
	jw.Uint16PtrValue(&v)
	if bf.String() != `65535` {
		t.Fatal(`Incorrect result for: 65535 !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

/**
 * uint32 is the set of all unsigned 32-bit integers. Range: 0 through 4294967295.
 */

func TestJsonWriter_Uint32(t *testing.T) { // {{{
	var v uint32
	bf := bytes.NewBuffer(make([]byte, 0, 10))
	jw := json.NewJsonWriter(bf)

	jw.Uint32Value(v)
	if bf.String() != `0` {
		t.Fatal(`Incorrect result for: 0 !=`, bf.String())
	}
	bf.Reset()

	v = 4294967295
	jw.Uint32Value(v)
	if bf.String() != `4294967295` {
		t.Fatal(`Incorrect result for: 4294967295 !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

func TestJsonWriter_Uint32Ptr(t *testing.T) { // {{{
	var v uint32
	bf := bytes.NewBuffer(make([]byte, 0, 10))
	jw := json.NewJsonWriter(bf)

	jw.Uint32PtrValue(&v)
	if bf.String() != `0` {
		t.Fatal(`Incorrect result for: 0 !=`, bf.String())
	}
	bf.Reset()

	v = 4294967295
	jw.Uint32PtrValue(&v)
	if bf.String() != `4294967295` {
		t.Fatal(`Incorrect result for: 4294967295 !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

/**
 * uint64 is the set of all unsigned 64-bit integers. Range: 0 through 18446744073709551615.
 */

func TestJsonWriter_Uint64(t *testing.T) { // {{{
	if strconv.IntSize != 64 {
		t.Skip(`Your system is not support 64-bit`)
	}

	var v uint64
	bf := bytes.NewBuffer(make([]byte, 0, 20))
	jw := json.NewJsonWriter(bf)

	jw.Uint64Value(v)
	if bf.String() != `0` {
		t.Fatal(`Incorrect result for: 0 !=`, bf.String())
	}
	bf.Reset()

	v = 18446744073709551615
	jw.Uint64Value(v)
	if bf.String() != `18446744073709551615` {
		t.Fatal(`Incorrect result for: 18446744073709551615 !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

func TestJsonWriter_Uint64Ptr(t *testing.T) { // {{{
	if strconv.IntSize != 64 {
		t.Skip(`Your system is not support 64-bit`)
	}

	var v uint64
	bf := bytes.NewBuffer(make([]byte, 0, 20))
	jw := json.NewJsonWriter(bf)

	jw.Uint64PtrValue(&v)
	if bf.String() != `0` {
		t.Fatal(`Incorrect result for: 0 !=`, bf.String())
	}
	bf.Reset()

	v = 18446744073709551615
	jw.Uint64PtrValue(&v)
	if bf.String() != `18446744073709551615` {
		t.Fatal(`Incorrect result for: 18446744073709551615 !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

/**
 * uint is an unsigned integer type that is at least 32 bits in size. It is a distinct type, however, and not an alias for, say, uint32.
 */

func TestJsonWriter_Uint(t *testing.T) { // {{{
	var v uint
	if strconv.IntSize < 64 { // 32-bit
		bf := bytes.NewBuffer(make([]byte, 0, 10))
		jw := json.NewJsonWriter(bf)

		jw.UintValue(v)
		if bf.String() != `0` {
			t.Fatal(`Incorrect result for: 0 !=`, bf.String())
		}
		bf.Reset()

		v = 4294967295
		jw.UintValue(v)
		if bf.String() != `4294967295` {
			t.Fatal(`Incorrect result for: 4294967295 !=`, bf.String())
		}
		bf.Reset()

		jw.Release()
	} else { // 64-bit
		bf := bytes.NewBuffer(make([]byte, 0, 20))
		jw := json.NewJsonWriter(bf)

		jw.UintValue(v)
		if bf.String() != `0` {
			t.Fatal(`Incorrect result for: 0 !=`, bf.String())
		}
		bf.Reset()

		v = 18446744073709551615
		jw.UintValue(v)
		if bf.String() != `18446744073709551615` {
			t.Fatal(`Incorrect result for: 18446744073709551615 !=`, bf.String())
		}
		bf.Reset()

		jw.Release()
	}
} // }}}

func TestJsonWriter_UintPtr(t *testing.T) { // {{{
	var v uint
	if strconv.IntSize < 64 {
		bf := bytes.NewBuffer(make([]byte, 0, 10))
		jw := json.NewJsonWriter(bf)

		jw.UintPtrValue(&v)
		if bf.String() != `0` {
			t.Fatal(`Incorrect result for: 0 !=`, bf.String())
		}
		bf.Reset()

		v = 4294967295
		jw.UintPtrValue(&v)
		if bf.String() != `4294967295` {
			t.Fatal(`Incorrect result for: 4294967295 !=`, bf.String())
		}
		bf.Reset()

		jw.Release()
	} else {
		bf := bytes.NewBuffer(make([]byte, 0, 20))
		jw := json.NewJsonWriter(bf)

		jw.UintPtrValue(&v)
		if bf.String() != `0` {
			t.Fatal(`Incorrect result for: 0 !=`, bf.String())
		}
		bf.Reset()

		v = 18446744073709551615
		jw.UintPtrValue(&v)
		if bf.String() != `18446744073709551615` {
			t.Fatal(`Incorrect result for: 18446744073709551615 !=`, bf.String())
		}
		bf.Reset()

		jw.Release()
	}
} // }}}

// --[ /uint ]--

// --[ float ]--

/**
 * float32 is the set of all IEEE-754 32-bit floating-point numbers.
 */

func TestJsonWriter_Float32(t *testing.T) { // {{{
	var v float32
	bf := bytes.NewBuffer(make([]byte, 0, 20))
	jw := json.NewJsonWriter(bf)

	jw.Float32Value(v)
	if bf.String() != `0` {
		t.Fatal(`Incorrect result for: 0 !=`, bf.String())
	}
	bf.Reset()

	v = 8765.39
	jw.Float32Value(v)
	if bf.String() != `8765.39` {
		t.Fatal(`Incorrect result for: 8765.39 !=`, bf.String())
	}
	bf.Reset()

	v = 876.539
	jw.Float32Value(v)
	if bf.String() != `876.539` {
		t.Fatal(`Incorrect result for: 876.539 !=`, bf.String())
	}
	bf.Reset()

	v = -8765.39
	jw.Float32Value(v)
	if bf.String() != `-8765.39` {
		t.Fatal(`Incorrect result for: -8765.39 !=`, bf.String())
	}
	bf.Reset()

	v = -876.539
	jw.Float32Value(v)
	if bf.String() != `-876.539` {
		t.Fatal(`Incorrect result for: -876.539 !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

func TestJsonWriter_Float32Ptr(t *testing.T) { // {{{
	var v float32
	bf := bytes.NewBuffer(make([]byte, 0, 20))
	jw := json.NewJsonWriter(bf)

	jw.Float32PtrValue(&v)
	if bf.String() != `0` {
		t.Fatal(`Incorrect result for: 0 !=`, bf.String())
	}
	bf.Reset()

	v = 8765.39
	jw.Float32PtrValue(&v)
	if bf.String() != `8765.39` {
		t.Fatal(`Incorrect result for: 8765.39 !=`, bf.String())
	}
	bf.Reset()

	v = 876.539
	jw.Float32PtrValue(&v)
	if bf.String() != `876.539` {
		t.Fatal(`Incorrect result for: 876.539 !=`, bf.String())
	}
	bf.Reset()

	v = -8765.39
	jw.Float32PtrValue(&v)
	if bf.String() != `-8765.39` {
		t.Fatal(`Incorrect result for: -8765.39 !=`, bf.String())
	}
	bf.Reset()

	v = -876.539
	jw.Float32PtrValue(&v)
	if bf.String() != `-876.539` {
		t.Fatal(`Incorrect result for: -876.539 !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

/**
 * float64 is the set of all IEEE-754 64-bit floating-point numbers.
 */

func TestJsonWriter_Float64(t *testing.T) { // {{{
	if strconv.IntSize != 64 {
		t.Skip(`Your system is not support 64-bit`)
	}

	var v float64
	bf := bytes.NewBuffer(make([]byte, 0, 20))
	jw := json.NewJsonWriter(bf)

	jw.Float64Value(v)
	if bf.String() != `0` {
		t.Fatal(`Incorrect result for: 0 !=`, bf.String())
	}
	bf.Reset()

	v = 418765.392
	jw.Float64Value(v)
	if bf.String() != `418765.392` {
		t.Fatal(`Incorrect result for: 418765.392 !=`, bf.String())
	}
	bf.Reset()

	v = -418765.392
	jw.Float64Value(v)
	if bf.String() != `-418765.392` {
		t.Fatal(`Incorrect result for: -418765.392 !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

func TestJsonWriter_Float64Ptr(t *testing.T) { // {{{
	if strconv.IntSize != 64 {
		t.Skip(`Your system is not support 64-bit`)
	}

	var v float64
	bf := bytes.NewBuffer(make([]byte, 0, 20))
	jw := json.NewJsonWriter(bf)

	jw.Float64PtrValue(&v)
	if bf.String() != `0` {
		t.Fatal(`Incorrect result for: 0 !=`, bf.String())
	}
	bf.Reset()

	v = 418765.392
	jw.Float64PtrValue(&v)
	if bf.String() != `418765.392` {
		t.Fatal(`Incorrect result for: 418765.392 !=`, bf.String())
	}
	bf.Reset()

	v = -418765.392
	jw.Float64PtrValue(&v)
	if bf.String() != `-418765.392` {
		t.Fatal(`Incorrect result for: -418765.392 !=`, bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

// --[ /float ]--

func TestJsonWriter_Array(t *testing.T) { // {{{
	bf := bytes.NewBuffer(make([]byte, 0, 2))
	jw := json.NewJsonWriter(bf)

	jw.ArrOpen().ArrClose()
	if bf.String() != "[]" {
		t.Fatal("Incorrect result for: [] !=", bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}

func TestJsonWriter_Object(t *testing.T) { // {{{
	bf := bytes.NewBuffer(make([]byte, 0, 2))
	jw := json.NewJsonWriter(bf)

	jw.ObjOpen().ObjClose()
	if bf.String() != "{}" {
		t.Fatal("Incorrect result for: {} !=", bf.String())
	}
	bf.Reset()

	jw.Release()
} // }}}
