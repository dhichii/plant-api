package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	return string(hash)
}

func MatchPassword(hashedPassword string, password string) bool {
	byteHash := []byte(hashedPassword)
	if err := bcrypt.CompareHashAndPassword(byteHash, []byte(password)); err != nil {
		return false
	}
	return true
}
