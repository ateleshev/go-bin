// go-bin: github.com/ateleshev/go-bin
package main

import (
	"io"
)

type reader struct {
	name string
}

func (v *reader) Open(n string) {
	v.name = n
}

func (v *reader) Close() error {
	println(v.name)
	return nil
}

// Ptr(s)

func example0() { // {{{
	c := &reader{}
	c.Open("file1")
	defer func() {
		c := c
		c.Close()
	}()

	c.Open("file2")
	defer func() {
		c := c
		c.Close()
	}()
} // }}}

func example1() { // {{{
	c := &reader{}
	c.Open("file1")
	defer c.Close()

	c.Open("file2")
	defer c.Close()
} // }}}

func example2() { // {{{
	c := &reader{}
	c.Open("file1")
	defer c.Close()

	c = &reader{}
	c.Open("file2")
	defer c.Close()
} // }}}

func example3() { // {{{
	c := &reader{}
	c.Open("file1")
	defer func() {
		c.Close()
	}()

	c.Open("file2")
	defer func() {
		c.Close()
	}()
} // }}}

func example4() { // {{{
	c := &reader{}
	c.Open("file1")
	defer func() {
		c.Close()
	}()

	c = &reader{}
	c.Open("file2")
	defer func() {
		c.Close()
	}()
} // }}}

func example5() { // {{{
	c := &reader{}
	c.Open("file1")
	defer func(c io.Closer) {
		c.Close()
	}(c)

	c.Open("file2")
	defer func(c io.Closer) {
		c.Close()
	}(c)
} // }}}

func example6() { // {{{
	c := &reader{}
	c.Open("file1")
	defer func(c io.Closer) {
		c.Close()
	}(c)

	c = &reader{}
	c.Open("file2")
	defer func(c io.Closer) {
		c.Close()
	}(c)
} // }}}

func example7() { // {{{
	c := reader{}
	c.Open("file1")
	defer c.Close()

	c.Open("file2")
	defer c.Close()
} // }}}

func example8() { // {{{
	c := reader{}
	c.Open("file1")
	defer c.Close()

	c = reader{}
	c.Open("file2")
	defer c.Close()
} // }}}

func example9() { // {{{
	c := reader{}
	c.Open("file1")
	defer func() {
		c.Close()
	}()

	c.Open("file2")
	defer func() {
		c.Close()
	}()
} // }}}

func example10() { // {{{
	c := reader{}
	c.Open("file1")
	defer func() {
		c.Close()
	}()

	c = reader{}
	c.Open("file2")
	defer func() {
		c.Close()
	}()
} // }}}

func example11() { // {{{
	c := reader{}
	c.Open("file1")
	defer func(c reader) {
		c.Close()
	}(c)

	c.Open("file2")
	defer func(c reader) {
		c.Close()
	}(c)
} // }}}

func example12() { // {{{
	c := reader{}
	c.Open("file1")
	defer func(c reader) {
		c.Close()
	}(c)

	c = reader{}
	c.Open("file2")
	defer func(c reader) {
		c.Close()
	}(c)
} // }}}

func main() { // {{{
	println("example0:")
	example0()
	println("example1:")
	example1()
	println("example2:")
	example2()
	println("example3:")
	example3()
	println("example4:")
	example4()
	println("example5:")
	example5()
	println("example6:")
	example6()
	println("example7:")
	example7()
	println("example8:")
	example8()
	println("example9:")
	example9()
	println("example10:")
	example10()
	println("example11:")
	example11()
	println("example12:")
	example12()
} // }}}
