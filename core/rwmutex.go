package core

import "sync"

// One reader, one writer.
type RWMutex struct {
	N            int // Total number of elements in protected array.
	a, b         int // half-open interval locked for writing
	c, d         int // half-open interval locked for reading
	writingframe int // time stamp of data currently being written in [a, b[
	lastread     int // time stamp of data last read in [c, d[
	state        sync.Mutex // protects the internal state, used in cond.
	cond         sync.Cond // wait condition: read/write is safe
	readers []*RMutex
}

func NewRWMutex(N int) *RWMutex {
	m := new(RWMutex)
	m.N = N
	m.cond = *(sync.NewCond(&m.state))
	m.writingframe = -1 // nothing yet written
	m.lastread = -1     // nothing yet read
	return m
}

// ______________________________________________________ Write

// Lock for writing [start, stop[.
func (m *RWMutex) Lock(start, stop int) {

	if start > stop || start >= m.N || stop > m.N || start < 0 || stop < 0 {
		Panicf("rwmutex: lock: invalid arguments: start=%v, stop=%v, n=%v", start, stop, m.N)
	}

	m.cond.L.Lock()
	//Debug("WLock", start, stop)
	if start == 0 {
		m.writingframe++
		//Debug("W new frame, writingframe=", m.writingframe)
	}
	m.a, m.b = start, start // noting is being written while waiting
	for !m.canWLock(start, stop) {
		//Debug("Wlock: wait")
		m.cond.Wait()
	}
	m.a, m.b = start, stop // update lock the interval
	m.cond.L.Unlock()
	m.cond.Broadcast()
}

// Can m safely lock for writing [start, stop[ ?
// Not thread-safe, assumes state mutex is locked.
func (m *RWMutex) canWLock(a, b int) (ok bool) {
	c, d := m.c, m.d
	//reason := "?"
	//defer func() { Debug("canWlock: [", a, ",", b, "[, [", c, ",", d, "[", ok, reason) }()

	// intersection of read & write interval:
	ok = !intersects(a, b, c, d)
	if !ok {
		//reason = "intersects"
		return
	}
	// make sure we don't overwrite data that has not yet been read.
	if a >= d {
		if m.stampOf(a) != m.lastread { // time stamp should be OK
			//reason = fmt.Sprint("stampOf", a, "==", m.stampOf(a), "!=", m.lastread)
			return false
		}
	}
	//reason = "ok"
	return true
}

// ______________________________________________________ Read

// Lock for reading [start, stop[.
func (m *RWMutex) rLock(start, stop int) {
	if start > stop || start >= m.N || stop > m.N || start < 0 || stop < 0 {
		Panicf("rwmutex: rlock: invalid arguments: start=%v, stop=%v, n=%v", start, stop, m.N)
	}

	m.cond.L.Lock()
	//Debug("RLock", start, stop)
	m.c, m.d = start, start
	for !m.canRLock(start, stop) {
		//Debug("Rlock: wait")
		m.cond.Wait()
	}
	m.c, m.d = start, stop
	if stop == m.N {
		m.lastread++
		//Debug("R new frame, lastread=", m.lastread)
	}
	m.cond.L.Unlock()
	m.cond.Broadcast() // TODO: benchmark with broadcast in/out lock.
}

// Can m safely lock for reading [start, stop[ ?
// Not thread-safe, assumes state mutex is locked.
func (m *RWMutex) canRLock(c, d int) (ok bool) {
	a, b := m.a, m.b
	//reason := "?"
	//defer func() { Debug("canRlock: [", a, ",", b, "[, [", c, ",", d, "[", ok, reason) }()

	ok = !intersects(a, b, c, d) // intersection should be empty
	if !ok {
		//reason = "intersects"
		return
	}
	// make sure we don't read data that has not yet been written.
	if c >= b {
		if m.stampOf(d) != m.lastread+1 { // time stamp should be OK
			//reason = fmt.Sprint("stampOf", d, "==", m.stampOf(d), "!=", m.lastread, "+ 1")
			ok = false
			return
		}
	}
	//reason = "ok"
	return true
}

// [a, b[ intersects [c, d[ ?
func intersects(a, b, c, d int) bool {
	return max(a, c) < min(b, d)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Time stamp when data at index has last been written.
// Not thread-safe, assumes state mutex is locked.
func (m *RWMutex) stampOf(index int) int {
	if index < m.a {
		return m.writingframe
	}
	if index >= m.b {
		return m.writingframe - 1
	}
	Panicf("rwmutex: writingframe: invalid index: start=%v, stop=%v, index=%v", m.a, m.b, index)
	return -2 // silence gc (dummy value)
}
