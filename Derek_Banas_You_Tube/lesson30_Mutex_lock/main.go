package main

import "sync"

type Account struct {
	balance int
	lock    sync.Mutex
	// A Mutex is a mutual exclusion lock.
	// The zero value for a Mutex is an unlocked mutex.
	//
	// A Mutex must not be copied after first use.
	//
	// In the terminology of the Go memory model,
	// the n'th call to Unlock “synchronizes before” the m'th call to Lock
	// for any n < m.
	// A successful call to TryLock is equivalent to a call to Lock.
	// A failed call to TryLock does not establish any “synchronizes before”
	// relation at all.
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func main() {

}
