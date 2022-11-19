package str

import "math/rand"

var (
	listGenerateType = map[string]string{
		"small":    "abcdefghijklmnopqrstuvwxyz",
		"big":      "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"num":      "1234567890",
		"symbol":   "!@#$%^&*()-_+=",
		"alphanum": "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
	}
)

func GenerateString(length int, generateTypes ...string) string {
	defaultLetter := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if len(generateTypes) != 0 {
		defaultLetter = ""
		for _, generateType := range generateTypes {
			defaultLetter += listGenerateType[generateType]
		}
	}

	b := make([]byte, length)

	for i := range b {
		b[i] = defaultLetter[rand.Intn(len(defaultLetter))]
	}

	return string(b)
}
