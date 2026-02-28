package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	AccountNumber string
	Balance       float64
	OwnerName     string
	Mutex         sync.Mutex
}

func (acc *Account) Deposit(amount float64) error {
	acc.Mutex.Lock()
	defer acc.Mutex.Unlock()
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be greater than zero")
	}
	acc.Balance += amount
	fmt.Printf("Deposited $%.2f into account %s. New Balance is $%.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil
}

func (acc *Account) Withdraw(amount float64) error {
	acc.Mutex.Lock()
	defer acc.Mutex.Unlock()
	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be greater than zero")
	}
	if acc.Balance < amount {
		return fmt.Errorf("insufficient balance. attempt to withdraw $%.2f from account %s with current balance of $%.2f", amount, acc.AccountNumber, acc.Balance)
	}
	acc.Balance -= amount
	fmt.Printf("Withdrawn $%.2f from account %s. New Balance is $%.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil
}

func (acc *Account) GetBalance() float64 {
	acc.Mutex.Lock()
	defer acc.Mutex.Unlock()
	return acc.Balance
}

func (acc *Account) String() string {
	acc.Mutex.Lock()
	defer acc.Mutex.Unlock()
	return fmt.Sprintf("Account Number [%s] \nOwner: %s, \nBalance: $%.2f", acc.AccountNumber, acc.OwnerName, acc.Balance)
}

type SavingsAccount struct {
	Account
	InterestRate float64
}

func (sa *SavingsAccount) AddInterest() {
	interest := sa.Balance * sa.InterestRate
	fmt.Printf("Adding interest $%.2f to savings account %s.\n", interest, sa.AccountNumber)
	err := sa.Deposit(interest)
	if err != nil {
		fmt.Printf("AddInterest: Error depositing $%.2f to savings account %s: %v\n", interest, sa.AccountNumber, err)
	}
}

type OverdraftAccount struct {
	Account
	OverdraftLimit float64
}

func (oa *OverdraftAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be greater than zero")
	}
	if (oa.Balance + oa.OverdraftLimit) < amount {
		return fmt.Errorf("insufficient balance. attempt to withdraw $%.2f from account %s exceeds overdraft limit of $%.2f with current balance of $%.2f", amount, oa.AccountNumber, oa.Balance+oa.OverdraftLimit, oa.Balance)
	}
	oa.Balance -= amount
	fmt.Printf("Withdrawn $%.2f from overdraft account %s. New Balance is $%.2f\n", amount, oa.AccountNumber, oa.Balance)
	return nil
}

func main() {
	fmt.Println("============== Bank Account System==============  ")
	var wg sync.WaitGroup

	savAcc := &SavingsAccount{
		Account: Account{
			AccountNumber: "SAV001",
			Balance:       1000,
			OwnerName:     "Alice Wonderland",
		},
		InterestRate: 0.02,
	}
	fmt.Println("\n==============Savings Account Operations==============")
	fmt.Println(savAcc.Account.String())

	depositAmount := 500.00
	// withdraw := 200.00

	for i := range 5 {
		wg.Add(1)
		go func(deposit float64) {
			defer wg.Done()
			time.Sleep(time.Duration(deposit) * time.Microsecond)
			err := savAcc.Deposit(deposit)
			if err != nil {
				fmt.Printf("Error depositing $%.2f to savings account %s: %v\n", deposit, savAcc.AccountNumber, err)
			}
		}(depositAmount + float64(i*100))
	}

	wg.Wait()
	// err := savAcc.Deposit(depositAmount)
	// if err != nil {
	// 	fmt.Printf("Error depositing $%.2f to savings account %s: %v\n", depositAmount, savAcc.AccountNumber, err)
	// }

	// savAcc.AddInterest()

	// err := savAcc.Withdraw(withdraw)
	// if err != nil {
	// 	fmt.Printf("Error withdrawing $%.2f from savings account %s: %v\n", withdraw, savAcc.AccountNumber, err)
	// }

	// fmt.Println("Final savings details:\n", savAcc.Account.String())

	// overdAcc := OverdraftAccount{
	// 	Account: Account{
	// 		AccountNumber: "OverD002",
	// 		Balance:       100,
	// 		OwnerName:     "Bob Spenderland",
	// 	},
	// 	OverdraftLimit: 200,
	// }

	// fmt.Println("\n========= Overdraft Account Operations ========= ")
	// fmt.Println(overdAcc.Account.String())

	// err = overdAcc.Deposit(depositAmount)
	// if err != nil {
	// 	fmt.Printf("Error depositing $%.2f to overdraft account %s: %v\n", depositAmount, overdAcc.AccountNumber, err)
	// }

	// err = overdAcc.Withdraw(700)
	// if err != nil {
	// 	fmt.Printf("Error withdrawing $%.2f from overdraft account %s: %v\n", withdraw, overdAcc.AccountNumber, err)
	// }
	// err = overdAcc.Withdraw(110)
	// if err != nil {
	// 	fmt.Printf("Error withdrawing $110 from overdraft account %s: %v\n", overdAcc.AccountNumber, err)
	// }

	// fmt.Println("Final overdraft details:\n", overdAcc.Account.String())
}
