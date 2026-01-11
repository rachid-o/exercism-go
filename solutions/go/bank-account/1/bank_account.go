package account

import "sync"

type Account struct {
	mu     sync.Mutex
	Bal    int64
	IsOpen bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{Bal: amount, IsOpen: true}
}

func (a *Account) Balance() (int64, bool) {
	return a.Bal, a.IsOpen
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	balance, isOpen := a.Balance()
	if !isOpen || (amount+balance) < 0 {
		return 0, false
	}
	a.Bal += amount
	return a.Bal, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	balanceBeforeClose := a.Bal
	if a.IsOpen == false {
		return 0, false
	}
	a.IsOpen = false
	a.Bal = 0
	return balanceBeforeClose, true
}
