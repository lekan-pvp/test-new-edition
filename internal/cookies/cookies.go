package cookies

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strings"
)

var key = []byte("secret key")

func New() *http.Cookie {

	id := uuid.NewString()

	h := hmac.New(sha256.New, key)
	h.Write([]byte(id))
	dst := h.Sum(nil)

	cookie := &http.Cookie{
		Name:  "token",
		Value: fmt.Sprintf("%s:%x", id, dst),
		Path:  "/",
	}
	return cookie
}

func CheckCookie(cookie *http.Cookie) bool {
	values := strings.Split(cookie.Value, ":")

	data, err := hex.DecodeString(values[1])
	if err != nil {
		log.Fatal(err)
		return false
	}

	id := values[0]

	log.Println("IN CheckCookie:", id)

	h := hmac.New(sha256.New, key)
	h.Write([]byte(id))
	sign := h.Sum(nil)

	return hmac.Equal(sign, data)
}
