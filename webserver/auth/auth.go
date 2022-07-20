package auth

import (
	"gateway/package/log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

const (
	accessTokenCookieName  = "access-token"
	refreshTokenCookieName = "refresh-token"

	//TODO (jwm) get the top secret from configuration file
	jwtSecretKey        = "0ifp5bw(!1j8bq#2bwd24)bn!0$gco6hhoce^!7tmprdaf$1z7" //TODO store this into the configuration
	jwtRefreshSecretKey = "0ifp5bw(!1j7bq#2bwd24)bn!0$gco5hhoce^!7tmprdaf$1z7" //TODO store this into the configuration

	expirationTime = 30 //experation time in minutes
	AuthContextKey = "auth_user"
)

var (
	accessTokenMiddleware  echo.MiddlewareFunc
	refreshTokenMiddleware echo.HandlerFunc
)

//GetJWTSecret returns secreat key
func GetJWTSecret() string {
	return jwtSecretKey
}

func GetRefreshJWTSecret() string {
	return jwtRefreshSecretKey
}

func GenerateTokensAndSetCookies(id int32, email string, c echo.Context) error {

	accessToken, exp, err := generateAccessToken(id, email)
	if err != nil {
		return err
	}

	setTokenCookie(accessTokenCookieName, accessToken, exp, c)
	setUserCookie(email, exp, c)
	refreshToken, exp, err := generateRefreshToken(id, email)
	if err != nil {
		return err
	}
	setTokenCookie(refreshTokenCookieName, refreshToken, exp, c)

	return nil
}

func ClearSession(c echo.Context) {
	log.Infoln("clearing user session")
	clearUserCookie(c)
	clearTokenCookie(accessTokenCookieName, c)
	clearTokenCookie(refreshTokenCookieName, c)
}

func generateAccessToken(id int32, email string) (string, time.Time, error) {
	// Declare the expiration time of the token
	expTime := time.Now().Add(expirationTime * time.Minute)

	return generateToken(id, email, expTime, []byte(GetJWTSecret()))
}

func generateRefreshToken(id int32, email string) (string, time.Time, error) {
	// Declare the expiration time of the token
	expTime := time.Now().Add(expirationTime * 10 * time.Minute) //TODO (jwm): adjust the refresh time

	return generateToken(id, email, expTime, []byte(GetRefreshJWTSecret()))
}

func generateToken(id int32, email string, expTime time.Time, secret []byte) (string, time.Time, error) {
	// Create the JWT claims, which includes the username and expiry time
	claims := &JWTCustomClaims{
		ID:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expTime.Unix(),
			//ExpiresAt: 0,
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expTime, nil
}

func setTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	cookie.HttpOnly = true

	c.SetCookie(cookie)
}

func setUserCookie(email string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = email
	cookie.Expires = expiration
	cookie.Path = "/"
	c.SetCookie(cookie)
}

func clearTokenCookie(name string, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = ""
	cookie.MaxAge = -1
	cookie.Path = "/"
	c.SetCookie(cookie)
}
func clearUserCookie(c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = ""
	cookie.MaxAge = -1
	cookie.Path = "/"
	c.SetCookie(cookie)
}

// TokenRefresherMiddleware middleware, which refreshes JWT tokens if the access token is about to expire.
func TokenRefresherMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	refreshTokenMiddleware = func(c echo.Context) error {

		if SkipperLoginCheck(c) {
			return next(c)
		}

		// If the user is not authenticated (no user token data in the context), don't do anything.
		if c.Get("user") == nil {
			return next(c)
		}
		// Gets user token from the context.
		u := c.Get("user").(*jwt.Token)

		claims := u.Claims.(*JWTCustomClaims)

		// We ensure that a new token is not issued until enough time has elapsed
		// In this case, a new token will only be issued if the old token is within
		// expirationTime/2 mins of expiry.
		if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < expirationTime/2*time.Minute {
			// Gets the refresh token from the cookie.
			rc, err := c.Cookie(refreshTokenCookieName)
			if err == nil && rc != nil {
				// Parses token and checks if it valid.
				tkn, err := jwt.ParseWithClaims(rc.Value, claims, func(token *jwt.Token) (interface{}, error) {
					return []byte(GetRefreshJWTSecret()), nil
				})
				if err != nil {
					if err == jwt.ErrSignatureInvalid {
						c.Response().Writer.WriteHeader(http.StatusUnauthorized)
					}
				}

				if tkn != nil && tkn.Valid {
					// If everything is good, update tokens.
					_ = GenerateTokensAndSetCookies(claims.ID, claims.Email, c)
				}
			}
		}

		return next(c)
	}
	return refreshTokenMiddleware
}

func GetUserFromContext(c echo.Context) (int32, string) {
	user := c.Get("user")
	if user == nil {
		return 0, ""
	}

	uu := user.(*jwt.Token)
	claims := uu.Claims.(*JWTCustomClaims)
	if claims == nil {
		return 0, ""
	}
	return claims.ID, claims.Email
}
