package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, counter, 3)
	})
	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		// charged wg with count
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				// count inside wg -1
				wg.Done()
			}()
		}
		// block until count = 0
		wg.Wait()

		assertCount(t, counter, wantedCount)
	})
}

func assertCount(t testing.TB, got *Counter, want int) {
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
