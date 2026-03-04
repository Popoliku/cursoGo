package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	c := Counter{}
	wantedCount := 1000
	var wg sync.WaitGroup
	wg.Add(wantedCount)

	for i := 0; i < wantedCount; i++ {
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()

	if c.Value() != wantedCount {
		t.Errorf("got %d, want %d", c.Value(), wantedCount)
	}
}
