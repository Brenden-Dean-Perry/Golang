package concurrnecy

// Go routines are lightweight threads managed by the Go runtime.
// They are cheaper than traditional threads and can be created in large numbers.
// They enable concurrent programming, allowing multiple functions to run simultaneously.
// Note concurrency is not parallelism; concurrency is about dealing with lots of things at once,
// while parallelism is about doing lots of things at once. However, parrellism is a subset of concurrency.

// Go routines will achieve parallelism if there are multiple CPU cores available and the Go scheduler
// decides to run them in parallel. The Go runtime manages the scheduling of goroutines onto OS threads,
// which can be distributed across multiple CPU cores.

import (
	"fmt"
	"sync"
	"time"
)

var dbData = []string{"data1", "data2", "data3", "data4", "data5"}

// fetchData simulates fetching data from a database with a delay
// Launch each DB call in a separate goroutine. Currently, it will spawn 5 goroutines in the background but then exit the function becuase it does not wait for them to complete.
func SimulateFetchData(query string) {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		go simulateDBCall(i) // Launches in background but does not wait for completion
	}
	fmt.Println("Fetch time:", time.Since(t0))
}

func simulateDBCall(i int) {
	time.Sleep(2 * time.Second) // Simulate delay
	fmt.Printf("Result for %s\n", dbData[i])
}

// In order to wait for all goroutines to finish before exiting the function, we can use sync.WaitGroup.
var wg = sync.WaitGroup{}

// fetchData simulates fetching data from a database with a delay
func SimulateFetchDataWithWaitGroup(query string) {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)                         // Increment the WaitGroup counter
		go simulateDBCallWithWaitGroup(i) // Launches in background
	}
	wg.Wait() // Wait for all goroutines to finish and counter to reach zero
	fmt.Println("Fetch time:", time.Since(t0))
}

func simulateDBCallWithWaitGroup(i int) {
	time.Sleep(2 * time.Second) // Simulate delay
	fmt.Printf("Result for %s", dbData[i])
	wg.Done() // Decrement the counter when the goroutine completes
}

// What if we want to save the results from each goroutine?
// Well if we try to append to a slice from multiple goroutines, we will likely run into race conditions (unexpected results as a result of concurrent writes).
// In order to avoid this, we can use a mutex (mutual exclusion lock) to ensure that only one goroutine can access the shared resource at a time.
var results = []string{}
var mu = sync.Mutex{}

// fetchData simulates fetching data from a database with a delay
func SimulateFetchDataWithWaitGroupAndMutexLock(query string) {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)                                 // Increment the WaitGroup counter
		go simulateDBCallWithWaitGroupAndMutex(i) // Launches in background and uses a mutex to protect shared resource
	}
	wg.Wait() // Wait for all goroutines to finish and counter to reach zero
	fmt.Println("Fetch time:", time.Since(t0))
}

// simulateDBCallWithWaitGroupAndMutex simulates a database call and appends the result to a shared slice
// Note: Where you place the lock and unlock is important. You want to minimize the time the lock is held to avoid blocking other goroutines!
// If you place the lock before the sleep, it will serialize the calls and defeat the purpose of using goroutines.
func simulateDBCallWithWaitGroupAndMutex(i int) {
	time.Sleep(2 * time.Second) // Simulate delay
	mu.Lock()
	results = append(results, fmt.Sprintf("Result for %s", dbData[i]))
	mu.Unlock()
	wg.Done() // Decrement the counter when the goroutine completes
}

// Note: A generic mutex will block both reads and writes of all goroutines while one goroutine holds the lock.
// You may want this. You may not.
// If you have a lot of reads and few writes, you can use a RWMutex (read-write mutex) which allows multiple readers but only one writer at a time.
var rwMu = sync.RWMutex{}

// simulateDBCallWithWaitGroupAndMutex simulates a database call and appends the result to a shared slice
// Note: Where you place the lock and unlock is important. You want to minimize the time the lock is held to avoid blocking other goroutines!
// If you place the lock before the sleep, it will serialize the calls and defeat the purpose of using goroutines.
func SimulateDBCallWithWaitGroupAndRWMutex(i int) {
	time.Sleep(2 * time.Second) // Simulate delay
	rwMu.Lock()
	results = append(results, fmt.Sprintf("Result for %s", dbData[i]))
	rwMu.Unlock()
	logWithRWMutex()
	wg.Done() // Decrement the counter when the goroutine completes
}

// The RWMutex has two types of locks:
// The Lock and Unlock methods are used for writing (exclusive access).
// The RLock and RUnlock methods are used for reading (shared access).
func logWithRWMutex() {
	rwMu.RLock() // Acquire read lock
	fmt.Println("The current results are:", results)
	rwMu.RUnlock() // Ensure the read lock is released
}

// Warning! Note all locks should be released. If a lock is not released, it will cause a deadlock and the program will hang.
// Also, all locks must be clear before the goroutine will gain write access.
// But if there are reads can happen as long as there are no writes. Reads can happen concurrently but not in parallel! Reads will happen one at a time but only if there are no write locks.

// In summary, goroutines are lightweight threads that enable concurrent programming in Go.
// They are managed by the Go runtime and can be synchronized using WaitGroups and Mutexes to protect shared resources.
