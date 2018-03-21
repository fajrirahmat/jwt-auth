package util

import (
	"testing"
)

func TestValidateWithRSAPublicKey(t *testing.T) {
	//your token is here
	var tokenString = ""
	//your RSA public key is here
	var pemPublicKey = `
-----BEGIN PUBLIC KEY-----                                                                                                                                                                     
KEY                                                                                                                                                               
-----END PUBLIC KEY-----
`
	key := []byte(pemPublicKey)
	valid := ValidateWithRSAPublicKey(tokenString, key)
	if !valid {
		t.Error("Token is not valid")
	}
}
