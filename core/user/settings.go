package user

import (
	"os"
	"time"
	"gorm.io/gorm"

)

var (
	// Database Reference
	DB *gorm.DB

	// JWT Configuration
	JwtSecret      = []byte(getEnv("JWT_SECRET", "change-me-immediately"))
	TokenLifespan  = time.Hour * 24
	
	// Security Configuration
	BcryptCost     = 14
	AuthHeaderKey  = "Authorization"
	BearerPrefix   = "Bearer "
)

// Helper to handle env defaults
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}