package user

import (
	"fmt"
	"transactions/domain/account"
)


type User struct {
	ID string 
	CPF    string 
	Secret string 
	user_account account.Account
}
var user_cache = make(map[string]User)

func CreateUser(cpf string, secret string) User{
	if _, ok := user_cache[cpf]; ok {
		fmt.Print("User already exists")
		return user_cache[cpf]
	}

	new_user := User{
		CPF: cpf,
		Secret: secret,
		user_account: account.CreateAccount("José", cpf),
	}

	user_cache[new_user.CPF] = new_user
	return  new_user
}

