package vtgate

import (
	"testing"
	"time"

	"github.com/youtube/vitess/go/stats"
)

type DummyConnection struct {
	closed bool
}

func (dc *DummyConnection) Close() {
	dc.closed = true
}

func assertEqual(t *testing.T, name string, actual interface{}, expected interface{}) {
	if expected != actual {
		t.Errorf("wanted %v = %v, got %v", name, expected, actual)
	}
}

func assertSizeAndBusyCount(t *testing.T, c *Conntrack, expectedSize int, expectedBusyCount int) {
	assertEqual(t, "c.Size()", c.Size(), expectedSize)
	assertEqual(t, "c.busyCount", int(c.busyCount.Get()), expectedBusyCount)
}

func TestConntrackConnection(t *testing.T) {
	busyCount := stats.NewInt("TestConntrackConnection")
	c := NewConntrack(busyCount)

	dc := &DummyConnection{}

	assertSizeAndBusyCount(t, c, 0, 0)

	c.Update(dc, false)

	assertSizeAndBusyCount(t, c, 1, 0)

	// ensure we don't double count idle connections
	c.Update(dc, false)

	assertSizeAndBusyCount(t, c, 1, 0)

	// transition connection to busy
	c.Update(dc, true)

	assertSizeAndBusyCount(t, c, 1, 1)

	// ensure we don't double count busy connections
	c.Update(dc, true)

	assertSizeAndBusyCount(t, c, 1, 1)

	// ensure that a busy deleted connection updates both total and busy counts
	c.Delete(dc)

	assertSizeAndBusyCount(t, c, 0, 0)

	// bring back busy connection
	c.Update(dc, true)

	assertSizeAndBusyCount(t, c, 1, 1)

	// transition connection to idle
	c.Update(dc, false)

	assertSizeAndBusyCount(t, c, 1, 0)

	// ensure that an idle deleted connection just updates total count
	c.Delete(dc)

	assertSizeAndBusyCount(t, c, 0, 0)
}

func TestConntrackDrain(t *testing.T) {
	dc1 := &DummyConnection{}
	dc2 := &DummyConnection{}
	dc3 := &DummyConnection{}

	busyCount := stats.NewInt("TestConntrackDrain")
	c := NewConntrack(busyCount)

	assertSizeAndBusyCount(t, c, 0, 0)

	c.Update(dc1, true)
	c.Update(dc2, true)
	c.Update(dc3, true)

	assertSizeAndBusyCount(t, c, 3, 3)

	c.Update(dc1, false)

	assertSizeAndBusyCount(t, c, 3, 2)
	assertEqual(t, "dc1.closed", dc1.closed, false)
	assertEqual(t, "dc2.closed", dc2.closed, false)
	assertEqual(t, "dc3.closed", dc3.closed, false)

	c.Delete(dc2)

	assertSizeAndBusyCount(t, c, 2, 1)
	assertEqual(t, "dc1.closed", dc1.closed, false)
	assertEqual(t, "dc2.closed", dc2.closed, false)
	assertEqual(t, "dc3.closed", dc3.closed, false)

	drainCh := make(chan bool)
	go c.Drain(drainCh)

	// give some time for Drain() to do its thing
	time.Sleep(10 * time.Millisecond)

	// assert that idle connection was closed
	assertSizeAndBusyCount(t, c, 1, 1)
	assertEqual(t, "dc1.closed", dc1.closed, true)
	assertEqual(t, "dc2.closed", dc2.closed, false)
	assertEqual(t, "dc3.closed", dc3.closed, false)

	// transition busy connection to idle
	c.Update(dc3, false)
	assertSizeAndBusyCount(t, c, 1, 0)

	// wait for Drain() to complete
	<-drainCh

	assertSizeAndBusyCount(t, c, 0, 0)
	assertEqual(t, "dc1.closed", dc1.closed, true)
	assertEqual(t, "dc2.closed", dc2.closed, false) // not closed because we deleted it before Drain()
	assertEqual(t, "dc3.closed", dc3.closed, true)
}
