package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type claims struct {
	*jwt.StandardClaims
	Scope []string `json:"scope,omitonempty"`
	Aud   []string `json:"aud,omitonempty"`
}

//AuthHandler ... checks for auth header and get the token to then decode and verify
func AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			t := r.Header.Get("Authorization")

			if valid, err := validateAuthHeaderFormat(&t); !valid {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if valid, err := validateToken(t); !valid {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
}

//validateFormat ...
//arg t ... authoraization header value
//updates the argument to hold trimmed token value only (trims "Bearer")
func validateAuthHeaderFormat(t *string) (bool, error) {

	splitToken := strings.Split(*t, "Bearer")

	if len(splitToken) != 2 {
		return false, errors.New("Auth header is missing or not in a proper format")
	}

	*t = strings.TrimSpace(splitToken[1])
	return true, nil

}

//validateToken ...
func validateToken(t string) (bool, error) {

	//TODO: Get the key json web keys end point 
	publickey := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4ZA4i0K7AhdC/RdWoZ/H\nIDTy5Qvm4uRRp/5/8r3eXm4c5fgeuuAETqvlUUPH3D3xTHkAdd5xdW/zbafLG168\nFVPH7WbKhJXAM2D2kIG5IOOtpSC/lqLRXWSTk3tF3XyOoxgRbGC0cjr7msRB+sQx\npi5XO0Fk8TzaWrGDs1rHAtXub7m9R4L6X6mmo5o9E2gJVQHOB87Zte/NO5lB67Kk\nvlrsc593daUk8TykzuhASAHohq/6zO9U40JkB5gh2ZrOUkf2j+xlhZ8Jf8PfKa+8\nSvfhdXTYkeAx6jU2KkWs6537mQiIjWsrWrBNWtI/1ZjoJRJjzsUIXGgCk5HInNw0\nSQIDAQAB\n-----END PUBLIC KEY-----"

	token, err := jwt.ParseWithClaims(t, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM([]byte(publickey))
	})

	//not checking token.Valid yet coz of expiry
	if claims, ok := token.Claims.(*claims); ok {
		fmt.Printf("%v %v", claims.Scope, claims.StandardClaims.ExpiresAt)
		return true, nil
	} else {
		return false, err
	}

}

func (c claims) Valid() error {

	valid := true

	//this role is hard coded for this sample
	//TODO: make this configurable 
	if !contains(c.Scope, "app1.readonly") {
		valid = false
	}

	if valid {
		return nil
	} else {
		return errors.New("not a valid token")
	}

}

func contains(s []string, key string) bool {
	for _, v := range s {
		if v == key {
			return true
		}
	}
	return false
}
