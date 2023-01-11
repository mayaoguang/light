package jwt

import (
	"testing"
)

func TestCreateJWTToken(t *testing.T) {
	token, err := CreateJWTToken(map[string]interface{}{"user_id": "123456"})
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}

func TestParseToken(t *testing.T) {
	userId, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTIzNDU2In0.ZXppwyOTgnGRb0eqKgqz9fQ25br53udppW_vIeSjqvY")
	if err != nil {
		t.Error(err)
	}
	t.Log(userId)
}
