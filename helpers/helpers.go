package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"learngo/utils/faulterr"
	"math/rand"
	"strconv"
	"time"
	"unsafe"

	"github.com/gofrs/uuid"
)

const (
	letterBytes   = "0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func StringToInt64(str string) (int64, error) {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func GetMd5(input string) string {
	password_hash := md5.New()
	defer password_hash.Reset()
	password_hash.Write([]byte(input))
	return hex.EncodeToString(password_hash.Sum(nil))
}

func ValidateTokenExpiry(expiresAt time.Time) error {
	if expiresAt.UTC().Unix() < time.Now().UTC().Unix() {
		return nil
	}
	return nil
}

func GenerateRandomString(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())

	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func GenerateUID() (*uuid.UUID, *faulterr.FaultErr) {
	uid, uidErr := uuid.NewV4()
	if uidErr != nil {
		return nil, faulterr.NewInternalServerError(uidErr.Error())
	}

	return &uid, nil
}
