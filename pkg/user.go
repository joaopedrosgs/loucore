package pkg

import (
	"context"
	"github.com/joaopedrosgs/loucore/ent"
	"github.com/joaopedrosgs/loucore/ent/user"
	"golang.org/x/crypto/bcrypt"
)

func Authentication(email string, password string) (*ent.User, error) {
	u, err := client.User.Query().Where(user.EmailEQ(email)).Only(context.Background())
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))

	return u, err
}

func CreateAccount(name string, email string, password string) (*ent.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return nil, err
	}
	return client.User.Create().SetPasswordHash(string(passwordHash)).SetEmail(email).SetName(name).Save(context.Background())
}

func DeleteAccount(userID int) error {
	return client.User.DeleteOneID(userID).Exec(context.Background())
}
