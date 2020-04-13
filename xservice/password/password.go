package password

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(pwd string) (string, error) {
	//Generates a new hash from the given password
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	//Checks if the hash is correct for the given password
	return string(hash), bcrypt.CompareHashAndPassword(hash, []byte(pwd))
}

func Check(hash, pwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd)) == nil
}
