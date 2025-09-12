package general

import (
	"crypto/rand"
	"fmt"
	"time"
)

// 生成用户的UID
func GenerateUID() (string, error) {
	ts := time.Now().UnixMilli()

	var b [1]byte
	_, err := rand.Read(b[:])
	if err != nil {
		return "", err
	}
	r := int(b[0]) % 100

	return fmt.Sprintf("U%d%02d", ts, r), nil
}
