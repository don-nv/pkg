package hasher

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

func (h *Hasher) MakeSHA256_FromString(source string) (string, error) {
	if source == "" {
		return "", errors.New("empty source string")
	}

	var (
		sourceBytes = []byte(source)
		hash        = sha256.Sum256(sourceBytes)
	)
	return fmt.Sprintf("%x", hash), nil
}
