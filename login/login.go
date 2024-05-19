package login

//username and password description.
var userNames = []string{"admin", "user"}
var userPws = []string{"adminpw11", "userpw22"}

//login function.
func Login(username, password string) bool {
	//check usernames and passwords.
	for i := range userNames {
		if userNames[i] == username && userPws[i] == password {
			return true
		}
	}
	return false
}
