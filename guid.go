package goguid

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

const shortForm = "20060102150405.999"
const nByte = 256

//Main
func CreateToken() string {
	rnd := generateRnd(nByte)
	curTime := encodeTime(getTime(shortForm))
	return curTime + "__" + rnd
}

func CreateCustomToken(byteLen int, timeFormat string) string {
	rnd := generateRnd(byteLen)
	curTime := encodeTime(getTime(timeFormat))
	return curTime + "__" + rnd
}

//Return current time in const shortForm format
func getTime(format string) string {
	t := time.Now().UTC()
	return t.Format(format)
}

//Encode UTC Time to a base64(URL Friendly) string
func encodeTime(curTime string) string {
	return base64.URLEncoding.EncodeToString([]byte(curTime))
}

//Generate random string
func generateRnd(byteLen int) string {
	rnd, _ := GenerateRandomString(byteLen)
	return base64.URLEncoding.EncodeToString([]byte(rnd))
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_.,"
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}
