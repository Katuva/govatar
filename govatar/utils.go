package govatar

import (
	"crypto/sha256"
	"encoding/hex"
)

// SHA256Hash Hash a value using SHA256
func SHA256Hash(value string) string {
	h := sha256.New()
	h.Write([]byte(value))
	return hex.EncodeToString(h.Sum(nil))
}
