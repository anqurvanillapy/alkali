package lockpool_test

import (
	"fmt"
	"sync"
	"time"

	"github.com/anqurvanillapy/alkali/sync/lockpool"
)

func ExampleNewPool() {
	p := lockpool.NewPool(64, 128)

	var wg sync.WaitGroup
	wg.Add(4)

	op1 := func(l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()

		fmt.Println("op1")
	}
	op2 := func(l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()

		time.Sleep(500 * time.Millisecond)
		fmt.Println("op2")
	}
	op3 := func(l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()

		fmt.Println("op3")
	}
	op4 := func(l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()

		fmt.Println("op4")
	}

	// Serially initialize the tasks.
	l1 := p.New(1)
	l2 := p.New(2, 3)
	l3 := p.New(3, 3+64) // only locks 3rd slot
	l4 := p.New(3, 4)

	// Concurrently run the tasks.
	go op4(l4) // op2 wakes up op3, not op4
	go op3(l3) // waits for op2
	go op2(l2)
	go op1(l1) // no one locks it

	wg.Wait()

	// Output:
	// op1
	// op2
	// op3
	// op4
}
