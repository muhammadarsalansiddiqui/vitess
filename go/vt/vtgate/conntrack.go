package vtgate

import (
	"sync"

	"github.com/youtube/vitess/go/stats"
)

// Closer is an interface which provides a Close() function.
// We're not using io.Closer because that Close() returns error.
type Closer interface {
	Close()
}

// Conntrack tracks connections
type Conntrack struct {
	mu        *sync.Mutex
	busy      map[Closer]bool
	draining  bool
	cond      *sync.Cond
	busyCount *stats.Int
}

// NewConntrack creates a new Conntrack instance
func NewConntrack(busyCount *stats.Int) *Conntrack {
	mu := &sync.Mutex{}
	return &Conntrack{
		mu:        mu,
		busy:      make(map[Closer]bool),
		cond:      sync.NewCond(mu),
		busyCount: busyCount,
	}
}

// Update whether or not a connection is busy
func (c *Conntrack) Update(conn Closer, isBusy bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	oldIsBusy := c.busy[conn]
	if isBusy && !oldIsBusy {
		c.busyCount.Add(1)
	} else if oldIsBusy && !isBusy {
		c.busyCount.Add(-1)
	}
	c.busy[conn] = isBusy
	if !isBusy && c.draining {
		c.cond.Broadcast()
	}
}

// Delete a connection
func (c *Conntrack) Delete(conn Closer) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.busy[conn] {
		c.busyCount.Add(-1)
	}
	delete(c.busy, conn)
	if c.draining {
		c.cond.Broadcast()
	}
}

// Size returns the total number of connections
func (c *Conntrack) Size() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.busy)
}

// Drain closes connections as they become idle. It blocks until all connections have closed.
func (c *Conntrack) Drain(drainCh chan bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	defer close(drainCh)
	c.draining = true
	for finished := false; !finished; {
		finished = true
		for conn, connIsBusy := range c.busy {
			if !connIsBusy {
				conn.Close()
				delete(c.busy, conn)
			} else {
				finished = false
			}
		}
		if !finished {
			c.cond.Wait()
		}
	}
}
