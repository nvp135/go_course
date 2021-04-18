/*
	Data Race problem https://medium.com/german-gorelkin/race-8936927dba20
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	acc, acc1, acc2 := account{balance: 0}, account{balance: 0}, account{balance: 0}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(n int) {
			wrongDeposit(&acc, 1)
			right1Deposit(&acc1, 1)
			right2Deposit(&acc2, 1)
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Printf("balance =  %d\nbalance1 =  %d\nbalance2 =  %d\n", acc.balance, acc1.balance, acc2.balance)

	accFrom, accTo, accFrom1, accTo1 := account{balance: 10000}, account{balance: 0}, account{balance: 10000}, account{balance: 0}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(n int) {
			err := transfer1(&accFrom, &accTo, 1)
			if err != nil {
				fmt.Printf("error for n=%d\n", n)
			}
			err1 := transfer2(&accFrom1, &accTo1, 1)
			if err1 != nil {
				fmt.Printf("error for n=%d\n", n)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("accFrom.balance = %d\n accTo.balance = %d\n", accFrom.balance, accTo.balance)
	fmt.Printf("accFrom.balance = %d\n accTo.balance = %d\n", accFrom1.balance, accTo1.balance)
}

type account struct {
	balance int64
}

func wrongDeposit(acc *account, amount int64) {
	acc.balance += amount
}

var mu sync.Mutex

func right1Deposit(acc *account, amount int64) {
	mu.Lock()
	acc.balance += amount
	mu.Unlock()
}

func right2Deposit(acc *account, amount int64) {
	atomic.AddInt64(&acc.balance, amount)
}

func transfer1(accFrom, accTo *account, amount int64) error {
	if accFrom.balance < amount {
		return fmt.Errorf("accFrom.balance<amount")
	}
	accTo.balance += amount
	accFrom.balance -= amount
	return nil
}

func transfer2(accFrom, accTo *account, amount int64) error {
	mu.Lock()
	bal := accFrom.balance
	mu.Unlock()

	if bal < amount {
		return fmt.Errorf("accFrom.balance<amount")
	}

	mu.Lock()
	accTo.balance += amount
	mu.Unlock()

	mu.Lock()
	accFrom.balance -= amount
	mu.Unlock()

	return nil
}
