package account

import (
	"fmt"
	"time"
)

type Account struct {
	ID string 
	Name  string 
	Cpf    string 
	Secret string 
	Balance  int 
	Created_at  time.Time 
}

var account_cache = make(map[string]Account)

func CreateAccount(name string, cpf string) Account{
	if _, ok := account_cache[cpf]; ok {
		fmt.Print("Account already exists")
		return account_cache[cpf]
	}
	new_account := Account{
		Name: name,
		Cpf: cpf,
		Created_at: time.Now(),
		Balance: 0,
		ID: "1",
		Secret: "123",
	}

	account_cache[new_account.Cpf] = new_account
	return  new_account
}

func GetAccount(cpf string) Account{
	if result, ok := account_cache[cpf]; ok {
		return result
	}

	fmt.Print("Account not found")
	
	return Account{}
	
}

func GetBalance(id string) int{
	if result, ok := account_cache[id]; ok {
		return result.Balance
	}

	fmt.Print("Account not found")
	
	return 0
}

