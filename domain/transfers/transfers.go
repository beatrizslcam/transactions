package transfers

import (
	"fmt"
	"time"
	"transactions/domain/account"
)

type Transfer interface{
	CreateTransfer(account_origin_id string, account_destination_id string, amount int) Transfers

}


type Transfers struct {
	Id string 
	AccountOriginID string 
	Account_destination_id string 
	Amount int
	Created_at  time.Time 
}


func CreateTransfer(account_origin_id string, account_destination_id string, amount int) Transfers  {
	if validateAccount(account_origin_id) && validateAccount(account_destination_id) {
		if account.GetBalance(account_origin_id) > amount {
			return doTransfer(account_origin_id, account_destination_id, amount)
		}
		fmt.Println("Insufficient funds")
		return Transfers{}
}

fmt.Println("Account not found")

return Transfers{}
}

func validateAccount(account_origin_id string) bool {
	result := account.GetAccount(account_origin_id)
	return result != (account.Account{})
}

func doTransfer(account_origin_id string, account_destination_id string, amount int) Transfers {
	
	account_origin := account.GetAccount(account_origin_id)
	account_destination := account.GetAccount(account_destination_id)


	account.Deposit(account_destination.ID, amount)
	account.Withdraw(account_origin.ID, amount)




	return Transfers{
		AccountOriginID: account_origin_id,
		Account_destination_id: account_destination_id,
		Amount: amount,
		Created_at: time.Now(),
	}
}