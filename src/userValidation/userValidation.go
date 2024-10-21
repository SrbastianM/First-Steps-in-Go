package uservalidation

func ValidateUser(user, pass string) bool {
	if user == "pedro" && pass == "1234" {
		return true
	} else {
		return false
	}
}
