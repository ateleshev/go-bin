package logger

import (
	"fmt"
	"github.com/ateleshev/go-bin/errutil"
	"sync"
)

const (
	ErrLoggerNotIdentified = errutil.Error("logger not identified")
)

// size - size of buffer, quantity of messages can be buffered in memory before send to logger
func NewBufferedLogger(logger Logger, size int) Logger { // {{{
	return &bufferedLogger{
		logger:  logger,
		size:    size,
		stop:    make(chan interface{}, 0),
		chanels: make(map[Mode]msgChan, ModesQnt),
		wg:      new(sync.WaitGroup),
	}
} // }}}

type bufferedLogger struct {
	logger  Logger
	size    int
	stop    chan interface{}
	chanels map[Mode]msgChan
	wg      *sync.WaitGroup
	closed  bool
}

func (this *bufferedLogger) Name() string { // {{{
	if this.logger != nil {
		return this.logger.Name()
	}

	return ""
} // }}}

func (this *bufferedLogger) Mode() Mode { // {{{
	if this.logger != nil {
		return this.logger.Mode()
	}

	return ModeNone
} // }}}

func (this *bufferedLogger) CheckMode(m Mode) bool { // {{{
	if this.logger != nil {
		return this.logger.CheckMode(m)
	}

	return false
} // }}}

func (this *bufferedLogger) Size() int { // {{{
	return this.size
} // }}}

func (this *bufferedLogger) Reset() { // {{{
	if this.logger != nil {
		this.logger.Reset()
		this.logger = nil
	}

	this.size = 0
	this.closed = false
	for m := range this.chanels {
		close(this.chanels[m])
		delete(this.chanels, m)
	}
} // }}}

func (this *bufferedLogger) msgHandler(ch msgChan, f msgWriter) { // {{{
	defer this.wg.Done()

	for {
		select {
		case m := <-ch:
			this.write(f, m)
			continue
		case <-this.stop:
			for len(ch) > 0 {
				this.write(f, <-ch)
			}
			return
		}
	}
} // }}}

func (this *bufferedLogger) write(f msgWriter, m *msg) { // {{{
	f(this.logger, m)
	m.Release()
} // }}}

func (this *bufferedLogger) Open() (err error) { // {{{
	if this.logger == nil {
		err = ErrLoggerNotIdentified
		return
	}

	if err = this.logger.Open(); err != nil {
		return
	}

	for mode, _ := range ModeNames {
		if this.logger.CheckMode(mode) {
			this.chanels[mode] = make(msgChan, this.size)

			this.wg.Add(1)
			go this.msgHandler(this.chanels[mode], getMsgWriter(mode))
		}
	}

	return
} // }}}

func (this *bufferedLogger) Close() (err error) { // {{{
	this.closed = true

	// Stop all handlers
	close(this.stop)

	// Wait stops
	this.wg.Wait()

	// Close logger
	err = this.logger.Close()

	return
} // }}}

func (this *bufferedLogger) Panic(v ...interface{}) { // {{{
	panic(fmt.Sprint(v...))
} // }}}

func (this *bufferedLogger) Panicf(f string, v ...interface{}) { // {{{
	panic(fmt.Sprintf(f, v...))
} // }}}

func (this *bufferedLogger) Info(v ...interface{}) { // {{{
	if this.closed {
		return
	}

	if _, ok := this.chanels[ModeInfo]; ok {
		this.chanels[ModeInfo] <- newMsg(v...)
	}
} // }}}

func (this *bufferedLogger) Infof(f string, v ...interface{}) { // {{{
	if this.closed {
		return
	}

	if _, ok := this.chanels[ModeInfo]; ok {
		this.chanels[ModeInfo] <- newFmtMsg(f, v...)
	}
} // }}}

func (this *bufferedLogger) Error(v ...interface{}) { // {{{
	if this.closed {
		return
	}

	if _, ok := this.chanels[ModeError]; ok {
		this.chanels[ModeError] <- newMsg(v...)
	}
} // }}}

func (this *bufferedLogger) Errorf(f string, v ...interface{}) { // {{{
	if this.closed {
		return
	}

	if _, ok := this.chanels[ModeError]; ok {
		this.chanels[ModeError] <- newFmtMsg(f, v...)
	}
} // }}}

func (this *bufferedLogger) Debug(v ...interface{}) { // {{{
	if this.closed {
		return
	}

	if _, ok := this.chanels[ModeDebug]; ok {
		this.chanels[ModeDebug] <- newMsg(v...)
	}
} // }}}

func (this *bufferedLogger) Debugf(f string, v ...interface{}) { // {{{
	if this.closed {
		return
	}

	if _, ok := this.chanels[ModeDebug]; ok {
		this.chanels[ModeDebug] <- newFmtMsg(f, v...)
	}
} // }}}

func (this *bufferedLogger) Access(v ...interface{}) { // {{{
	if this.closed {
		return
	}

	if _, ok := this.chanels[ModeAccess]; ok {
		this.chanels[ModeAccess] <- newMsg(v...)
	}
} // }}}

func (this *bufferedLogger) Accessf(f string, v ...interface{}) { // {{{
	if this.closed {
		return
	}

	if _, ok := this.chanels[ModeAccess]; ok {
		this.chanels[ModeAccess] <- newFmtMsg(f, v...)
	}
} // }}}
