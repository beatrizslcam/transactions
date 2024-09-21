package account_test

import (
	"testing"
	"transactions/domain/account"
)

func TestCreateAccount(t *testing.T) {
	account := account.CreateAccount("José", "123456789")
	if account.Name != "José" {
		t.Errorf("Expected José, got %s", account.Name)
	}
}	

