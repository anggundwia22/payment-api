package services

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your_secret_key") // Ganti dengan kunci rahasia Anda

// GenerateToken membuat token JWT untuk user berdasarkan ID pelanggan.
func GenerateToken(customerID int) (string, error) {
	// Membuat klaim
	claims := jwt.MapClaims{
		"customer_id": customerID,
		"exp":         time.Now().Add(24 * time.Hour).Unix(), // Token berlaku 24 jam
		"iat":         time.Now().Unix(),                    // Waktu token dibuat
	}

	// Membuat token dengan klaim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Menandatangani token dengan secret key
	return token.SignedString(jwtSecret)
}

// ValidateToken memvalidasi token JWT dan mengembalikan klaim.
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	// Parsing token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Memastikan token menggunakan metode HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Mengambil klaim dari token jika valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
