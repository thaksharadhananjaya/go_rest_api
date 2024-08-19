package middleware

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var publicKey *rsa.PublicKey

func init() {
	// Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found or failed to load .env file")
    }

    var err error
    publicKeyBytes, err := ioutil.ReadFile(os.Getenv("JWT_PUBLIC_KEY_PATH")) // Path to your public key
    if err != nil {
        log.Fatalf("Error reading public key file: %v", err)
    }
    publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
    if err != nil {
        log.Fatalf("Error parsing public key: %v", err)
    }
}

func JWTMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized request"})
            c.Abort()
            return
        }

        // Extract the Bearer token
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized request"})
            c.Abort()
            return
        }

        tokenString := parts[1]

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
                return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
            }
            return publicKey, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized request"})
            c.Abort()
            return
        }

        c.Next()
    }
}
