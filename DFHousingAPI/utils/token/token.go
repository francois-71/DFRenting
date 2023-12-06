package token

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(user_id string, role string) (string, error) {

	token_lifespan,err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return "",err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	claims["role"] = role
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))

}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (string, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, ok := claims["user_id"].(string)
		if !ok {
			return "", fmt.Errorf("User ID is not a string")
		}
		return uid, nil
	}
	return "", nil
}



func ExtractTokenRole(c *gin.Context) (string, error) {
    tokenString := ExtractToken(c)
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(os.Getenv("API_SECRET")), nil
    })
    if err != nil {
        return "", err
    }
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return "", fmt.Errorf("Invalid token")
    }
	// print all claims
	fmt.Printf("claims: %v", claims)
    // Retrieve the role claim from the token's claims
    roleClaim, exists := claims["role"]
    if !exists {
        return "", fmt.Errorf("Role claim does not exist")
    }

    // Convert the role claim to a string
    role, ok := roleClaim.(string)
    if !ok {
        return "", fmt.Errorf("Role claim is not a string")
    }

    // Ensure the role claim matches expected values
    if role != "admin" && role != "user" {
        return "", fmt.Errorf("Invalid role value")
    }

    return role, nil
}


