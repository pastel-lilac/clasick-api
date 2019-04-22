package hash

import "golang.org/x/crypto/bcrypt"

type hashUtil struct {}

type IHashUtil interface {
	CreateHash(pass string) (string, error)
	CheckHash(hash, pass string) error
}

func NewHashUtil() IHashUtil {
	return &hashUtil{}
}

func (util *hashUtil) CreateHash(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func (util *hashUtil) CheckHash(hash, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
}