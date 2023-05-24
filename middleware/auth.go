package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	flexcreek "github.com/ekholme/flex_creek"
	"github.com/golang-jwt/jwt/v5"
)

var jwt_key = os.Getenv("JWT_KEY")

// defining my claims
type CustomClaims struct {
	Username string `json:"username"`
	UserID   string `json:"userid"`
	jwt.RegisteredClaims
}

// auth struct
type Auth struct {
	Claims  *CustomClaims `json:"claims"`
	Token   string        `json:"token"`
	Expires *time.Time    `json:"-"`
}

// creating a new Auth object
// this will get called at login, I think
func CreateAuth(u *flexcreek.User) *Auth {
	exp := time.Now().Add(2 * time.Hour)

	claims := &CustomClaims{
		Username: u.Username,
		UserID:   u.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "flexcreek",
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	return &Auth{
		Claims:  claims,
		Expires: &exp,
	}
}

// generate a jwt
// this will also get called when logging in
func GenerateToken(a *Auth) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.Claims)

	tokenStr, err := token.SignedString([]byte(jwt_key))

	if err != nil {
		return err
	}

	a.Token = tokenStr

	return nil
}

// validate the jwt
func ValidateJWT(tokenStr string) (*jwt.Token, error) {
	tkn, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwt_key), nil
	})

	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, errors.New("token not Valid")
	}

	return tkn, nil
}

// RESUME WITH ACTUAL MIDDLEWARE WRAPPER
func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("FLEXAUTH")

		if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				writeJSON(w, http.StatusBadRequest, "cookie not found")
			default:
				writeJSON(w, http.StatusInternalServerError, "something went wrong")
			}
			return
		}

		token, err := ValidateJWT(cookie.Value)

		if err != nil {
			writeJSON(w, http.StatusUnauthorized, err)
			return
		}

		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "flexclaims", claims)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}

// dumb to redefine this but w/e
func writeJSON(w http.ResponseWriter, statusCode int, v any) {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(v)

}
