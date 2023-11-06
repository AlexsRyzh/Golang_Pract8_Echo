package utils

import (
	"github.com/gorilla/securecookie"
	"log"
	"net/http"
	"os"
	"time"
)

func CreateCookie(name, value string) *http.Cookie {

	hashKey := []byte(os.Getenv("HASH_COOKIE_KEY"))
	blockKey := []byte(os.Getenv("BLOCK_COOKIE_KEY"))
	var s = securecookie.New(hashKey, blockKey)

	encoded, err := s.Encode(name, value)
	if err != nil {
		log.Fatal(err)
	}
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = encoded
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Secure = true
	cookie.Path = "/cookies"
	return cookie
}

func GetCookie(name, value string) (*string, error) {
	hashKey := []byte(os.Getenv("HASH_COOKIE_KEY"))
	blockKey := []byte(os.Getenv("BLOCK_COOKIE_KEY"))
	var s = securecookie.New(hashKey, blockKey)

	var decoded string
	err := s.Decode(name, value, &decoded)
	if err != nil {
		return nil, err
	}

	return &decoded, nil

}
