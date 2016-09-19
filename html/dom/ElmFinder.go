package dom

import "io"
import "io/ioutil"
import "strings"
import "bytes"
import "errors"
import "unicode/utf8"

var (
	ElmFinderBufferCap = 256
	ElmFinderResultCap = 512

	ErrEmptyFindPath = errors.New("Empty find path")
	ErrIsNotHtml     = errors.New("Is not HTML")
)

type ElmFinder struct {
	reader      io.Reader
	buffer      []byte
	result      []byte
	maxLoadSize int64
	loaded      int64
}

func NewElmFinder(r io.Reader) *ElmFinder { // {{{
	f := getElmFinder()
	f.reader = r
	return f
} // }}}

func (this *ElmFinder) SetMaxLoadSize(v int64) { // {{{
	this.maxLoadSize = v
} // }}}

func (this *ElmFinder) NextLoadSize() int64 { // {{{
	if this.maxLoadSize < 0 { // No limit to load
		return int64(ElmFinderBufferCap)
	} else if this.maxLoadSize == 0 || this.loaded >= this.maxLoadSize {
		return int64(0)
	}

	if (this.maxLoadSize - this.loaded) < int64(ElmFinderBufferCap) {
		return int64(this.maxLoadSize - this.loaded)
	}

	return int64(ElmFinderBufferCap)
} // }}}

func (this *ElmFinder) Find(path string) ([]byte, error) { // {{{
	var err error
	var found bool
	var i, n, it, size int
	var loadSize int64

	if len(path) == 0 {
		return nil, ErrEmptyFindPath
	}
	elms := strings.Split(path, ".")
	for {
		it++
		if loadSize = this.NextLoadSize(); loadSize == 0 {
			goto end_find
		}

		if this.buffer, err = ioutil.ReadAll(io.LimitReader(this.reader, loadSize)); err != nil {
			if err == io.EOF {
				goto end_find
			}
			return nil, err
		}

		this.loaded += int64(len(this.buffer))

		// Check is html received, only on first iteration
		if it == 1 {
			for i < len(this.buffer) {
				switch this.buffer[i] {
				case ' ', '\t', '\n', '\r':
					i++
					continue
				case '<':
					if i > 0 {
						this.buffer = this.buffer[i:]
						i = 0
					}
					goto analyze_html
				default:
					return nil, ErrIsNotHtml
				}
			}
		}

	analyze_html:
		this.result = append(this.result, this.buffer[i:]...)

		for i < len(this.result) {
			switch this.result[i] {
			case utf8.RuneSelf:
				_, size = utf8.DecodeRune(this.result[i:])
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
