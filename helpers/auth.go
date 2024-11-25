package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/shaammiru/final-task-pbi-rakamin-fullstack-muhammad-syamil/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateToken(payload models.User) (string, error) {
	claims := UserClaims{
		ID:       payload.ID,
		Username: payload.Username,
		Email:    payload.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "task-5-pbi-fullstack-developer-muhammadsyamil",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GetTokenClaims(tokenString string) (UserClaims, error) {
	claims := UserClaims{}

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return claims, err
	}

	if !token.Valid {
		return claims, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

func GetClaimsFromContext(c *gin.Context) (UserClaims, error) {
	claims, exists := c.Get("claims")
	if !exists {
		return UserClaims{}, jwt.ErrTokenRequiredClaimMissing
	}

	return claims.(UserClaims), nil
}
