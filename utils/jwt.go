package utils

import (
	"go-rest-api/models"

	"github.com/go-chi/jwtauth/v5"
)

// privated
var _tokenAuth *jwtauth.JWTAuth

func InitTokenAuth() {
	_tokenAuth = jwtauth.New("HS256", []byte("login"), nil)
}
func TokenAuth() *jwtauth.JWTAuth {
	return _tokenAuth
}
func GenerateToken(model models.JwtModel) (string, error) {
	mapData, err := StructToMap(model)
	if err != nil {
		return "", err
	}
	_, token, err := _tokenAuth.Encode(mapData)
	if err != nil {
		return "", err
	}
	return token, err

}
