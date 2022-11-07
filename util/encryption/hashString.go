package encryption

import "golang.org/x/crypto/bcrypt"

func HashString(value string) (result string, e error) {
	byteValue := []byte(value)

	hash, err := bcrypt.GenerateFromPassword(byteValue, bcrypt.MinCost)

	if err != nil {
		e = err
		return
	}

	result = string(hash)
	return
}
