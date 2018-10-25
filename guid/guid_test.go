package guidutil

import (
	"log"
	"regexp"
	"testing"
)

func TestToken(t *testing.T) {
	token := CreateToken()
	if token == "" {
		t.Errorf("Missing token")
	}

	r, err := regexp.Compile("([a-zA-Z0-9-_=]+)__([a-zA-Z0-9-_=]+)")
	if err != nil {
		log.Print(token)
		t.Errorf("Regex Issues")
	}
	if !r.MatchString(token) {
		log.Print(token)
		t.Errorf("Malformed Token")
	}

}
