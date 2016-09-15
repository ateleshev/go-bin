package url

import "errors"
import "unicode/utf8"
import "github.com/ateleshev/go-bin/bytes/byteutil"

const (
	PathStartByte = '/'

	MinServerLen = 1                // a | localhost | localhost:8080 | 127.0.0.1:8080
	MinPathLen   = 1                // / | /path/to | /path/to/id/12345
	MinLen       = 7 + MinServerLen // --[ http://a ]--
)

var (
	httpScheme = []byte("http")

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
func FetchSchemeIndex(v []byte) (int, bool, error) { // {{{
	var i int
	var secure bool

	for i < len(httpScheme) {
		if v[i] == utf8.RuneSelf || byteutil.ToLower(v[i]) != httpScheme[i] {
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
func FetchServerIndex(v []byte) (int, error) { // {{{
	var i int
	for i < len(v) {
		if byteutil.IsWordCharacter(v[i]) {
			i++
			continue
		}

		switch v[i] {
		case '.', ':', '@', '!', '~', '-':
			i++
			continue
		}

		goto end_server_index
	}

end_server_index:
	if i < MinServerLen {
		return 0, ErrIsNotUrlServer
	}

	return i, nil
} // }}}

/**
 * /some/path/123456
 */
func FetchPathIndex(v []byte) (int, error) { // {{{
	var i int

	if v[i] != PathStartByte {
		goto end_path_index
	}
	i++

	for i < len(v) {
		if byteutil.IsWordCharacter(v[i]) {
			i++
			continue
		}

		switch v[i] {
		case '.', ':', '@', '!', '~', '-', '/':
			i++
			continue
		}

		goto end_path_index
	}

end_path_index:
	if i < MinPathLen {
		return 0, ErrIsNotUrlPath
	}

	return i, nil
} // }}}

/**
 * ?param1=123&param2=abcd&form1[column1]=data
 */
func FetchQueryIndex(v []byte) (int, error) { // {{{
	return 0, ErrIsNotUrlQuery
} // }}}

/**
 * #param1=1;form[abc]=123;param2=test
 */
func FetchFragmentIndex(v []byte) (int, error) { // {{{
	return 0, ErrIsNotUrlFragment
} // }}}
