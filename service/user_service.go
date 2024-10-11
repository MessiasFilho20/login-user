package service

import (
	"context"
	"errors"
	"login-user/prisma/db"
	"login-user/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService struct {
	Client *db.PrismaClient
}

func (us *UserService) Register(ctx context.Context, input UserInput) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = us.Client.User.CreateOne(
		db.User.Email.Set(input.Email),
		db.User.Password.Set(string(hashedPassword)),
	).Exec(ctx)

	return err
}
func (us *UserService) Login(ctx context.Context, input UserInput) (string, error) {

	user, err := us.Client.User.FindUnique(db.User.Email.Equals(input.Email)).Exec(ctx)

	if err != nil {
		return "", errors.New("usuario não encontrado")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", errors.New("Senha incorreta")
	}

	token, err := utils.GenerateJWT(user.ID, user.Email)

	if err != nil {
		return "", err
	}
	return token, nil
}
