package internal

import (
	"github.com/go-akka/akka"
	"time"
)

type childNameReserved struct {
	stats string
}

var (
	_childNameReservedInstance = &childNameReserved{stats: "Name Reserved"}
)

func (p childNameReserved) String() string {
	return p.stats
}

func (p *childNameReserved) ChildNameReserved() {}

func (p *childNameReserved) ChildStats() {}

type ChildRestartStats struct {
	uid   int64
	child akka.InternalActorRef

	maxNrOfRetriesCount         int
	restartTimeWindowStartNanos int
}

func NewChildRestartStats(child akka.InternalActorRef, maxNrOfRetriesCount, restartTimeWindowStartNanos int) akka.ChildRestartStats {
	stats := &ChildRestartStats{
		child:                       child,
		uid:                         child.Path().Uid(),
		maxNrOfRetriesCount:         maxNrOfRetriesCount,
		restartTimeWindowStartNanos: restartTimeWindowStartNanos,
	}

	return stats
}

func (p *ChildRestartStats) RequestRestartPermission(maxNrOfRetries, withinTimeMilliseconds int) bool {
	if maxNrOfRetries == 0 {
		return false
	}

	retriesIsDefined := maxNrOfRetries > 0
	windowIsDefined := withinTimeMilliseconds > 0

	if retriesIsDefined && !windowIsDefined {
		p.maxNrOfRetriesCount += 1
		return p.maxNrOfRetriesCount <= maxNrOfRetries
	}

	if windowIsDefined {
		if !retriesIsDefined {
			maxNrOfRetries = 1
		}
		return p.RetriesInWindowOkay(maxNrOfRetries, withinTimeMilliseconds)
	}

	return true
}

func (p *ChildRestartStats) RetriesInWindowOkay(retries, window int) bool {
	retriesDone := p.maxNrOfRetriesCount + 1
	now := time.Now().Nanosecond()
	windowStart := 0

	if p.restartTimeWindowStartNanos == 0 {
		p.restartTimeWindowStartNanos = now
		windowStart = now
	} else {
		windowStart = p.restartTimeWindowStartNanos
	}

	insideWindow := (now - windowStart) <= window

	if insideWindow {
		p.maxNrOfRetriesCount = retriesDone
		return retriesDone <= retries
	}

	p.maxNrOfRetriesCount = 1
	p.restartTimeWindowStartNanos = now
	return true
}

func (p *ChildRestartStats) Child() akka.InternalActorRef {
	return p.child
}

func (c *ChildRestartStats) ChildRestartStats() {}

func (p *ChildRestartStats) ChildStats() {}
