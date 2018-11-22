package account

import "sync"

type Account struct{
	closed bool
	balance int
	sync.Mutex
}

func Open(amount int) *Account {
	account := Account{}
	if amount >= 0 {
		account = Account{
			balance: amount,
			closed: false,
		}
	} else {
		return nil
	}
	return &account
}

func (a *Account) Balance() (balance int, ok bool) {
	return a.balance, !a.closed
}

func (a *Account) Close() (balance int, ok bool) {
	a.Lock()
	defer a.Unlock()
	if !a.closed {
		balance = a.balance
		a.closed = true
		ok = true
	}

	return balance, ok
}

func (a *Account) Deposit(amount int) (balance int, ok bool) {
	if a.closed {
		return 0, ok
	}
	a.Lock()
	defer a.Unlock()

	if a.balance+amount >= 0 {
		a.balance += amount
		ok = true
	}

	return a.balance, ok
}
