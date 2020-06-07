package loucore

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"lou-core/ent"
	"lou-core/ent/user"
)

func Authentication(email string, password string) bool {
	var userInfo struct {
		Email         string
		PassworldHash string
	}
	err := client.User.Query().Where(user.EmailEQ(email)).Select(user.FieldEmail, user.FieldPasswordHash).Scan(context.Background(), &userInfo)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(userInfo.PassworldHash), []byte(password))

	return err == nil

}
func CreateAccount(name string, email string, password string) (*ent.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 20)
	if err != nil {
		return nil, err
	}

	return client.User.Create().SetPasswordHash(string(passwordHash)).SetEmail(email).SetName(name).Save(context.Background())

}
