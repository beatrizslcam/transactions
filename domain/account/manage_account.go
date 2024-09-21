package account



func Deposit(id string, amount int) Account{
	result := GetAccount(id);
	
		if result !=  (Account{}) {
	
			result.Balance = result.Balance + amount;
			account_cache[id] = result;
			return result
		}
	
		return Account{}
	}
	
	func Withdraw(id string, amount int) Account{
		result := GetAccount(id);
		
			if  result != (Account{}) {
		
				result.Balance = result.Balance - amount;
				account_cache[id] = result;
				return result
			}
			
			return Account{}
		}
	