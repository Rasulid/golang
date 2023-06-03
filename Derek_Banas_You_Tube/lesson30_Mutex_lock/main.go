package main

import (
	"fmt"
	"sync"
	"time"
)

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

func main() {
	var acca Account
	acca.balance = 100
	fmt.Println(acca.GetBalance())
	for i := 0; i < 12; i++ {
		go acca.Withdraw(10)
	}
	time.Sleep(2 * time.Second)
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) Withdraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		//log.Fatal("Not enough money in account")
		fmt.Println("Not enough money in account")
	} else {
		fmt.Printf("%d: money in account\n    %d: withdraw\n", a.balance, v)
		a.balance -= v
	}
}
