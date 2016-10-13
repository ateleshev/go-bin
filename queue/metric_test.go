package queue_test

import "testing"
import "github.com/ateleshev/go-bin/queue"

func TestMetric(t *testing.T) {
	var n int64
	m := queue.NewMetric("test")
	defer func() {
		t.Log("Number", n, m.Number())
		alloc, sys, total := m.MemUsage()
		t.Log("MemUsage", alloc, sys, total)
		hAlloc, hSys, hObjects := m.HeapUsage()
		t.Log("HeapUsage", hAlloc, hSys, hObjects)
		sInuse, sSys, otherSys := m.StackUsage()
		t.Log("StackUsage", sInuse, sSys, otherSys)
	}()

	m.Begin()
	n = m.Increase()
	t.Log(n)
	n = m.Increase()
	t.Log(n)
	n = m.Increase()
	t.Log(n)
	n = m.Decrease()
	t.Log(n)
	m.Finish()
}
