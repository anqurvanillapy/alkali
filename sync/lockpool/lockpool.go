package lockpool

import (
	"sync"
	"sync/atomic"
)

type Pool interface {
	// New creates a locker that guards a slice of IDs.
	New(ids ...uint64) sync.Locker
}

type (
	lockQueue = chan *locker
	lockSlots = []lockQueue
	lockPool  struct {
		poolSize int
		slots    lockSlots
	}
	locker struct {
		total, ready    uint32
		runCh, finishCh chan struct{}
	}
)

func NewPool(poolSize, queueSize int) Pool {
	slots := make(lockSlots, poolSize, poolSize)
	for i := range slots {
		q := make(lockQueue, queueSize)
		slots[i] = q
		go func() {
			for l := range q {
				if atomic.AddUint32(&l.ready, 1) == l.total {
					l.runCh <- struct{}{}
				}
				<-l.finishCh
			}
		}()
	}
	return &lockPool{poolSize: poolSize, slots: slots}
}

func (p *lockPool) New(ids ...uint64) sync.Locker {
	l := &locker{runCh: make(chan struct{}), finishCh: make(chan struct{})}

	var (
		slots   []int
		slotSet = make(map[int]struct{})
	)
	// Make a unique slice of slots to lock.
	for _, ino := range ids {
		slot := int(ino % uint64(p.poolSize))
		if _, ok := slotSet[slot]; ok {
			continue
		}
		slotSet[slot] = struct{}{}
		slots = append(slots, slot)
	}
	l.total = uint32(len(slots))

	for _, slot := range slots {
		p.slots[slot] <- l
	}

	return l
}

func (l *locker) Lock()   { <-l.runCh }
func (l *locker) Unlock() { close(l.finishCh) }
