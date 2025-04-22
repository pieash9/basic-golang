package auth

func extractSession() string {
	return "session logged in"
}

func GetSession() string {
	return extractSession()
}
