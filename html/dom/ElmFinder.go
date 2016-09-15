package dom

import "io"
import "io/ioutil"
import "strings"
import "bytes"
import "errors"
import "unicode/utf8"

var (
	ElmFinderResultCap = 1024
	ElmFinderReadLen   = int64(256)

	ErrEmptyFindPath = errors.New("Empty find path")
)

type ElmFinder struct {
	reader io.Reader
	result []byte
}

func NewElmFinder(r io.Reader) *ElmFinder { // {{{
	f := getElmFinder()
	f.reader = r
	return f
} // }}}

func (this *ElmFinder) Find(path string) ([]byte, error) { // {{{
	var i, n int
	var found bool

	if len(path) == 0 {
		return nil, ErrEmptyFindPath
	}
	elms := strings.Split(path, ".")
	for {
		buf, err := ioutil.ReadAll(io.LimitReader(this.reader, ElmFinderReadLen))
		if err != nil {
			if err == io.EOF {
				goto end_find
			}
			return nil, err
		}
		this.result = append(this.result, buf...)

		for i < len(this.result) {
			switch this.result[i] {
			case utf8.RuneSelf:
				_, size := utf8.DecodeRune(this.result[i:])
				i += size
				continue
			case '<':
				i++
				l := len(elms[n])
				if !found {
					if i+l <= len(this.result) && strings.ToLower(elms[n]) == string(bytes.ToLower(this.result[i:i+l])) {
						if n == len(elms)-1 {
							found = true
						} else {
							n++
						}
					}
					i += l
					continue
				}

				if this.result[i] == '/' && strings.ToLower(elms[n]) == string(bytes.ToLower(this.result[i+1:i+1+l])) {
					this.result = this.result[:i-1]
					goto end_find
				}
				continue
			case '>':
				this.result = this.result[i+1:]
				i = 0
				continue
			}
			i++
		}
	}

end_find:
	if !found {
		this.result = this.result[:0]
	}

	return this.result, nil
} // }}}

func (this *ElmFinder) Result() []byte { // {{{
	return this.result
} // }}}

func (this *ElmFinder) Reset() { // {{{
	this.reader = nil
	this.result = this.result[:0]
} // }}}

func (this *ElmFinder) Release() { // {{{
	putElmFinder(this)
} // }}}
