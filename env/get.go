package env

import (
	"os"
	"strconv"
)

// n - variable name
// d - default value
func GetBool(n string, d bool) bool { // {{{
	v := os.Getenv(n)
	if len(v) > 0 {
		if b, err := strconv.ParseBool(v); err == nil {
			return b
		}
	}

	return d
} // }}}

// n - variable name
// d - default value
func GetFloat64(n string, d float64) float64 { // {{{
	v := os.Getenv(n)
	if len(v) > 0 {
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
	}

	return d
} // }}}

// n - variable name
// d - default value
func GetFloat32(n string, d float32) float32 { // {{{
	return float32(GetFloat64(n, float64(d)))
} // }}}

// n - variable name
// d - default value
func GetInt64(n string, d int64) int64 { // {{{
	v := os.Getenv(n)
	if len(v) > 0 {
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return i
		}
	}

	return d
} // }}}

// n - variable name
// d - default value
func GetInt32(n string, d int32) int32 { // {{{
	return int32(GetInt64(n, int64(d)))
} // }}}

// n - variable name
// d - default value
func GetInt16(n string, d int16) int16 { // {{{
	return int16(GetInt64(n, int64(d)))
} // }}}

// n - variable name
// d - default value
func GetInt8(n string, d int8) int8 { // {{{
	return int8(GetInt64(n, int64(d)))
} // }}}

// n - variable name
// d - default value
func GetInt(n string, d int) int { // {{{
	return int(GetInt64(n, int64(d)))
} // }}}

// n - variable name
// d - default value
func GetUint64(n string, d uint64) uint64 { // {{{
	v := os.Getenv(n)
	if len(v) > 0 {
		if u, err := strconv.ParseUint(v, 10, 64); err == nil {
			return u
		}
	}

	return d
} // }}}

// n - variable name
// d - default value
func GetUint32(n string, d uint32) uint32 { // {{{
	return uint32(GetUint64(n, uint64(d)))
} // }}}

// n - variable name
// d - default value
func GetUint16(n string, d uint16) uint16 { // {{{
	return uint16(GetUint64(n, uint64(d)))
} // }}}

// n - variable name
// d - default value
func GetUint8(n string, d uint8) uint8 { // {{{
	return uint8(GetUint64(n, uint64(d)))
} // }}}

// n - variable name
// d - default value
func GetUint(n string, d uint) uint { // {{{
	return uint(GetUint64(n, uint64(d)))
} // }}}

// n - variable name
// d - default value
func GetString(n string, d string) string { // {{{
	v := os.Getenv(n)
	if len(v) > 0 {
		return v
	}

	return d
} // }}}
