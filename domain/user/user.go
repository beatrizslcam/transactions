package user

import (
	"transactions/domain/account"
)


type User struct {
	ID string 
	CPF    string 
	Secret string 
	user_account account.Account
}

