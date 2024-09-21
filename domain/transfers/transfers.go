package transfers

import (
	"time"
)


type Transfers struct {
	Id string 
	AccountOriginID string 
	Account_destination_id string 
	Amount int
	Created_at  time.Time 
}

