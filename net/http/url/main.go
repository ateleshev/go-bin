package url

import "errors"
import "unicode/utf8"
import "github.com/ateleshev/go-bin/bytes/byteutil"

const (
	MinLen = 8 // --[ http://a ]--
)

var (
	schemePrefix = []byte("http")

	ErrIsNotUrlScheme   = errors.New("Is not url scheme")
	ErrIsNotUrlServer   = errors.New("Is not url server")
	ErrIsNotUrlPath     = errors.New("Is not url path")
	ErrIsNotUrlQuery    = errors.New("Is not url query")
	ErrIsNotUrlFragment = errors.New("Is not url fragment")
)

/**
 * http:// | https://
 *
 * @return (index, secure, error)
 */
func SchemeIndex(v []byte) (int, bool, error) { // {{{
	var i int
	var secure bool

	for i < len(schemePrefix) {
		if v[i] == utf8.RuneSelf || byteutil.ToLower(v[i]) != schemePrefix[i] {
			return 0, false, ErrIsNotUrlScheme
		}
		i++
	}

	if byteutil.ToLower(v[i]) == 's' {
		secure = true
		i++
	}

	if string(v[i:i+3]) == "://" {
		return i + 3, secure, nil
	}

	return 0, false, ErrIsNotUrlScheme
} // }}}

/**
 * user:pass@host:port
 * user:pass@127.0.0.1:port
 */
func AuthorityIndex(v []byte) (int, error) { // {{{
	var i int
	for i < len(v) {
		if !byteutil.IsASCII(v[i]) {
			goto end_server_index
		}

		switch v[i] {
		case '/', '?', '#', ' ':
			goto end_server_index
		}

		i++
	}

end_server_index:
	if i < 1 {
		return 0, ErrIsNotUrlServer
	}

	return i, nil
} // }}}

/**
 * /some/path/123456
 */
func PathIndex(v []byte) (int, error) { // {{{
	var i, size int

	if v[i] != '/' {
		goto end_path_index
	}
	i++

	for i < len(v) {
		size = 1
		switch v[i] {
		case '?', '#', ' ':
			goto end_path_index
		}

		if v[i] == utf8.RuneSelf {
			_, size = utf8.DecodeRune(v[i:])
		}

		i += size
	}

end_path_index:
	if i < 1 {
		return 0, ErrIsNotUrlPath
	}

	return i, nil
} // }}}

/**
 * ?param1=123&param2=abcd&form1[column1]=data
 */
func QueryIndex(v []byte) (int, error) { // {{{
	var i, size int

	if v[i] != '?' {
		goto end_query_index
	}
	i++

	for i < len(v) {
		size = 1
		switch v[i] {
		case '#', ' ':
			goto end_query_index
		}

		if v[i] == utf8.RuneSelf {
			_, size = utf8.DecodeRune(v[i:])
		}

		i += size
	}

end_query_index:
	if i < 1 {
		return 0, ErrIsNotUrlQuery
	}

	return i, nil
} // }}}

/**
 * #param1=1;form[abc]=123;param2=test
 */
func FragmentIndex(v []byte) (int, error) { // {{{
	var i, size int

	if v[i] != '#' {
		goto end_fragment_index
	}
	i++

	for i < len(v) {
		size = 1
		switch v[i] {
		case ' ':
			goto end_fragment_index
		}

		if v[i] == utf8.RuneSelf {
			_, size = utf8.DecodeRune(v[i:])
		}

		i += size
	}

end_fragment_index:
	if i < 1 {
		return 0, ErrIsNotUrlFragment
	}

	return i, nil
} // }}}
