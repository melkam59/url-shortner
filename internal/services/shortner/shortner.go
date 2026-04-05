package services

import (
	"crypto/sha256"
	"encoding/base32"
)

type IShortnerService interface {
	Sha256(input string) []byte
	Encode(data []byte) string
}

type ShortnerService struct {
}

func NewShortnerService() IShortnerService {
	return &ShortnerService{}
}

func (s *ShortnerService) Sha256(input string) []byte {
	hash := sha256.New()
	hash.Write([]byte(input))

	return hash.Sum(nil)
}

func (s *ShortnerService) Encode(data []byte) string {
	dst := make([]byte, base32.StdEncoding.EncodedLen(len(data)))
	base32.StdEncoding.Encode(dst, data)

	return string(dst)

}
