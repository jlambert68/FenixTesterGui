package detailedExecutionsModel

import (
	"sync"
	"time"
)

// Refresh the TestCasesSummaryTable
func refreshTestCasesSummaryTable() {
	TestCasesSummaryTable.Refresh()
}

// refreshTestCasesSummaryTableThrottler controls the execution rate of a function, 'refreshTestCasesSummaryTable()'
type refreshTestCasesSummaryTableThrottler struct {
	Interval time.Duration
	ticker   *time.Ticker
	requests chan struct{}
	quit     chan struct{}
	wg       sync.WaitGroup
}

// newRefreshTestCasesSummaryTableThrottler creates a new refreshTestCasesSummaryTableThrottler
func newRefreshTestCasesSummaryTableThrottler(interval time.Duration) *refreshTestCasesSummaryTableThrottler {
	t := &refreshTestCasesSummaryTableThrottler{
		Interval: interval,
		ticker:   time.NewTicker(interval),
		requests: make(chan struct{}, 100), // Buffer up to 100 requests
		quit:     make(chan struct{}),
	}
	t.wg.Add(1)
	go t.run()
	return t
}

// run processes requests to execute 'refreshTestCasesSummaryTable' at a controlled rate
func (t *refreshTestCasesSummaryTableThrottler) run() {
	defer t.wg.Done()
	for {
		select {
		case <-t.ticker.C:
			select {
			case <-t.requests:
				refreshTestCasesSummaryTable()
			default:
				// No pending request, do nothing
			}
		case <-t.quit:
			return
		}
	}
}

// RequestRefreshTestCasesSummaryTable queues a request to execute the function
func (t *refreshTestCasesSummaryTableThrottler) RequestRefreshTestCasesSummaryTable() {

	// Check if the requests channel is empty before queuing a new request
	if len(t.requests) == 0 {
		select {
		case t.requests <- struct{}{}:
			// Put a request to refresh on the channel
			//fmt.Println("Request queued at", time.Now().Format(time.RFC3339))
		default:
			// When the queue reach maximum number of items, which can happen in extreme parallel execution mode, then it comes to this part
			//fmt.Println("Queue was not empty at", time.Now().Format(time.RFC3339))
		}
	} else {
		// equest skipped, queue not empty at
		//fmt.Println("Request skipped, queue not empty at", time.Now().Format(time.RFC3339))
	}
}

// Stop stops the throttler and waits for it to shut down cleanly
func (t *refreshTestCasesSummaryTableThrottler) Stop() {
	close(t.quit)
	t.wg.Wait()
	t.ticker.Stop()
}
