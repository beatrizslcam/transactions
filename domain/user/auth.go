package user

type Auth interface{
	IsAuthenticated(secret string, ID string) bool
	Authenticate(secret string, ID string) bool
}

func Authenticate(secret string, ID string) bool{
	return false
}

func IsAuthenticated(secret string, ID string) bool{
	return false
}




