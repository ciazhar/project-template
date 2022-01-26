package string

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/google/uuid"
	"strings"
)

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func ToArrayUUID(req string) []uuid.UUID {
	if req != "" {
		arr := strings.Split(req, ",")
		var res []uuid.UUID
		for i := range arr {
			res = append(res, uuid.MustParse(arr[i]))
		}
		return res
	} else {
		return nil
	}
}
