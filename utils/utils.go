package utils

import (
	"crypto/rand"
    "github.com/dgrijalva/jwt-go"
    "time"
	"github.com/gin-gonic/gin"
)

// ErrorResponse formats error responses uniformly
func ErrorResponse(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, gin.H{"error": message})
}

func GenerateSalt(length int) ([]byte, error) {
    salt := make([]byte, length)
    _, err := rand.Read(salt)
    if err != nil {
        return nil, err
    }
    return salt, nil
}

func GenerateJWTToken(email string) (string, error) {
    // Define the expiration time for the token (e.g., 1 hour)
    expirationTime := time.Now().Add(1 * time.Hour)

    // Create the JWT claims, including the email and expiration time
    claims := jwt.MapClaims{
        "email": email,
        "exp":   expirationTime.Unix(),
    }

    // Create the token with the claims and the signing method
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign the token with a secret key to create the complete, encoded token
    secretKey := []byte("your-secret-key") // Replace this with your actual secret key
    signedToken, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }

    return signedToken, nil
}
