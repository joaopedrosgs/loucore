package loucore

import "testing"

func TestAuthentication(t *testing.T) {
	err := Authentication(testEnv.accounts[0].Email, "password")
	if err != nil {
		t.Errorf("Failed to perform Authentication with user user=%v, email=%v and password=%v: %v", testEnv.accounts[0].Name, testEnv.accounts[0].Email, "password", err)
		return
	}

}
func TestAuthenticationShouldFailBecauseWrongPassword(t *testing.T) {
	var password = "password"
	var wrongPassword = "password2"
	err := Authentication(testEnv.accounts[0].Email, wrongPassword)
	if err == nil {
		t.Errorf("Failed to check password with user user=%v, email=%v and password=%v: %v", testEnv.accounts[0].Name, testEnv.accounts[0].Email, password, err)
		return
	}
}

func TestCreateAccount(t *testing.T) {
	var userName = "user3"
	var email = "email3@email.com"
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
	var userName = "user4"
	var email = "email4@email.com"
	var password = "password"
	user, _ := CreateAccount(userName, email, password)
	err := DeleteAccount(user.ID)
	if err != nil {
		t.Errorf("Failed to DeleteAccount with id=%v", user.ID)
	}

}
