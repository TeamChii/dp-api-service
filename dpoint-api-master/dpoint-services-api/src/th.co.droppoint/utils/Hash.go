package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPin(pin string) string {
	h := sha256.New()
	var salt = "5tyx5ILe81U6Hnq0McLlDCddhDLTiL60PypSY86pJOhbUCV1WQaLFplFR56R"
	h.Write([]byte(pin + salt))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}
