package password_hasher

import (
	"crypto/md5"
	"encoding/hex"
	"to-do-list/app/model"
)

func HashPassword(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func CheckPassword(user model.User, password string) bool {
	return user.Password == HashPassword(password)
}
