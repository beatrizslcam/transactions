package account_test

import (
	"testing"
	"transactions/domain/account"
	"trasactions/domain/manage_account"
)


func TestDeposit(t *testing.T) {
	account := account.CreateAccount("Josefina", "123456adc789")
	manage_account.Deposit(100)
	if account.Balance != 100 {
		t.Errorf("Expected 100, got %d", account.Balance)
	}
}

func TestWithdraw(t *testing.T) {
	account := account.CreateAccount("Josefina da Silva", "123a456adc789")
	manage_account.Deposit(100)
	manage_account.Withdraw(10)
	if account.Balance != 90 {
		t.Errorf("Expected 0, got %d", account.Balance)
	}
}