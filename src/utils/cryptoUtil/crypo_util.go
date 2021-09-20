package cryptoUtil

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5(password string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}
