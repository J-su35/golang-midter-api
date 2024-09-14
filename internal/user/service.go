package user

import (
	"errors"
	"log"
	"midterm-api/internal/auth"
	"midterm-api/internal/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


type Service struct {
	Repository Repository
	secret     string
}

func NewService(db *gorm.DB, secret string) Service {
	return Service{
		Repository: NewRepository(db),
		secret:     secret,
	}
}


func (service Service) Login(req model.RequestLogin) (string, error) {
    // TODO: Check username and password here
    user, err := service.Repository.FindOneByUsername(req.Username)
    if err != nil {
        return "", errors.New("Invalid user or password")
    }
    // req.Password // req password
    // user.Password // hashed password
    if ok := checkPasswordHash(req.Password, user.Password); !ok {
        return "", errors.New("Invalid user or password")
    }
    // TODO: Create token here
    token, err := auth.CreateToken(user.Username, service.secret)
    if err != nil {
        log.Println("Fail to create token")
        return "", errors.New("Something went wrong")
    }
    return token, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}