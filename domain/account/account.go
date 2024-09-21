package account

import (
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



