// go-bin: github.com/ateleshev/go-bin
package main

import ()

type reader struct {
	name string
}

func (v *reader) Open(n string) {
	v.name = n
}

func (v *reader) Close() {
	println(v.name)
}

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

func main() { // {{{
	println("example1:")
	example1()
	println("example2:")
	example2()
	println("example3:")
	example3()
	println("example4:")
	example4()
} // }}}
