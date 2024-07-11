package appCommon

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSequence(n int) string {
	b := make([]rune, n)

	seed1 := rand.NewSource(time.Now().UnixNano())
	rand1 := rand.New(seed1)
	for i := range b {
		b[i] = letters[rand1.Intn(999999)%len(letters)]
	}
	return string(b)
}

func GenSalt(len int) string {
	if len < 0 {
		len = 50
	}
	return randSequence(len)
}

type md5Hash struct {
}

func NewMd5Hash() *md5Hash {
	return &md5Hash{}
}

func (h *md5Hash) Hash(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
