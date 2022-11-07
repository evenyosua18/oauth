package encryption

import "golang.org/x/crypto/bcrypt"

func CompareHashString(value, hashedValue string) bool {
	byteValue := []byte(value)
	byteHashValue := []byte(hashedValue)

	if err := bcrypt.CompareHashAndPassword(byteHashValue, byteValue); err != nil {
		return false
	}
	return true
}
