package loucore

import "testing"

func TestAuthentication(t *testing.T) {
	var userName = "user"
	var email = "email@email.com"
	var password = "password"
	user, _ := CreateAccount(userName, email, password)
	err := Authentication(email, password)
	if err != nil {
		t.Errorf("Failed to perform Authentication with user user=%v, email=%v and password=%v: %v", userName, email, password, err)
		return
	}
	_ = DeleteAccount(user.ID)

}
func TestAuthenticationShouldFailBecauseWrongPassword(t *testing.T) {
	var userName = "user"
	var email = "email@email.com"
	var password = "password"
	var wrongPassword = "password2"
	user, _ := CreateAccount(userName, email, password)
	err := Authentication(email, wrongPassword)
	if err == nil {
		t.Errorf("Failed to check password with user user=%v, email=%v and password=%v: %v", userName, email, password, err)
		return
	}
	_ = DeleteAccount(user.ID)

}

func TestCreateAccount(t *testing.T) {
	var userName = "user"
	var email = "email@email.com"
	var password = "password"
	user, err := CreateAccount(userName, email, password)
	if err != nil {
		t.Errorf("Failed to CreateAccount with user user=%v, email=%v and password=%v", userName, email, password)
		return
	}
	if user == nil {
		t.Errorf("Failed silently creating user account")
		return

	}
	_ = DeleteAccount(user.ID)

}

func TestDeleteAccount(t *testing.T) {
	var userName = "user2"
	var email = "email2@email.com"
	var password = "password2"
	user, _ := CreateAccount(userName, email, password)
	err := DeleteAccount(user.ID)
	if err != nil {
		t.Errorf("Failed to DeleteAccount with id=%v", user.ID)
	}

}
