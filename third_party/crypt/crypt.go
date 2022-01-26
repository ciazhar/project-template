package crypt

import (
	"github.com/ciazhar/emobi-service/third_party/response"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

func DecryptPassword(password *string) error {
	//decrypt pass
	hash, _ := regexp.MatchString(`^\$2[ayb]\$.{56}$`, *password)
	if *password != "" && !hash {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), 14)
		if err != nil {
			return response.Error(err)
		}
		*password = string(hashedPassword)
	}
	return nil
}
