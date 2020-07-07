package loucore

import (
	"context"
	"github.com/joaopedrosgs/loucore/ent"
	"github.com/joaopedrosgs/loucore/ent/user"
	"golang.org/x/crypto/bcrypt"
)

func Authentication(email string, password string) error {
	u, err := client.User.Query().Where(user.EmailEQ(email)).Only(context.Background())
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))

	return err
}

func CreateAccount(name string, email string, password string) (*ent.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return nil, err
	}

	return client.User.Create().SetPasswordHash(string(passwordHash)).SetEmail(email).SetName(name).Save(context.Background())

}

func DeleteAccount(userId int) error {
	return client.User.DeleteOneID(userId).Exec(context.Background())

}
