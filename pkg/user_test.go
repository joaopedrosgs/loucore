package pkg

import (
	"testing"
)

func TestAuthentication(t *testing.T) {
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Simple case 1", args{"email@email.com", "password"}, false},
		{"Simple case 2", args{"email2@email.com", "password"}, false},
		{"Authentication should fail since the password is wrong", args{"email@email.com", "password2"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Authentication(tt.args.email, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("Authentication() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateAccount(t *testing.T) {
	type args struct {
		name     string
		email    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Simple case 1", args{"test9999", "test9999@email.com", "apwfgsd"}, false},
		{"Create account should fail because the name is already taken", args{"name", "test9999@email.com", "apwfgsd"}, true},
		{"Create account should fail because the email is already taken", args{"test", "email@email.com", "apwfgsd"}, true},
		{"Create account should fail because the password is too small", args{"test", "email@email.com", ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateAccount(tt.args.name, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				_ = DeleteAccount(got.ID)
			}
		})
	}
}

func TestDeleteAccount(t *testing.T) {
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Simple case", args{testEnv.accounts[2].ID}, false},
		{"Delete account should fail since the account doesn't exists", args{-1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteAccount(tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
