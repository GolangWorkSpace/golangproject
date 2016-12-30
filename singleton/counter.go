package statistic

import "sync"

type Counter struct {
	number int
	mtx    sync.RWMutex
}

func (c *Counter) Add(n int) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.number += n
}
func (c *Counter) Get() int {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.number
}
func loadNumber() (n int, err error) {
	//...
	return
}

var (
	counter *Counter
	once    sync.Once
)

func LoadCounter() *Counter {
	once.Do(func() {
		if n, err := loadNumber(); err != nil {
			// log error or panic
		} else {
			counter = &Counter{
				number: n,
			}
		}
	})
	return counter
}
