package collections

// A Monitor is a channel that can be used to synchronize goroutines.
type Monitor chan struct{}

// NewMonitor creates a new monitor.
func NewMonitor() Monitor { return make(Monitor) }

// Wait waits for a notification on the monitor. Returns false if the monitor is closed.
func (m Monitor) Wait() bool { _, ok := <-m; return ok }

// Notify notifies the goroutine waiting on the monitor, if any.
func (m Monitor) Notify() { m <- struct{}{} }

// Close closes the monitor and notifies all goroutines waiting on the monitor.
func (m Monitor) Close() { close(m) }
