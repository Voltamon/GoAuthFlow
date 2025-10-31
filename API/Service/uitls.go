package auth

import (
	"regexp"
	"time"

	"crypto/sha256"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

func GetCurrentTime() time.Time {
	return time.Now()
}

func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

func GenerateHash(email string, password string) string {
	data := []byte(email + password)
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

func ResponseError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}
