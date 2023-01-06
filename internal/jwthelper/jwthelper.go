package jwthelper

import (
	jwt "github.com/dgrijalva/jwt-go"
)

const TOKEN = "asdghjpi34585fgdfs"

type Token struct {
	jwt.StandardClaims
	UserID uint64 `json:"user_id,omitempty"`
	Role   string `json:"role,omitempty"`
	Email  string `json:"email,omitempty"`
}

func NewToken(userID uint64, email string, role string) (string, error) {

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), Token{
		UserID: userID,
		Email:  email,
		Role:   role,
	})

	t, err := token.SignedString([]byte(TOKEN))

	return t, err
}

func Parse(token string) (uint64, bool) {

	tk := Token{}

	t, err := jwt.ParseWithClaims(token, &tk, func(t *jwt.Token) (interface{}, error) {
		return []byte(TOKEN), nil
	})

	if err != nil || !t.Valid {
		return 0, false
	}

	return tk.UserID, true
}
