package shortcode

import (
	"crypto/rand"
	"math/big"
)

type Shortcode struct {
}

func NewShortcode() *Shortcode {
	return &Shortcode{}
}

func (s *Shortcode) GenerateCode(codeLen int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, codeLen)
	for i := 0; i < codeLen; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		result[i] = letters[num.Int64()]
	}
	return string(result), nil
}
