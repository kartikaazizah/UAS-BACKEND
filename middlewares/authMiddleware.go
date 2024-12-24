package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Middleware to check user role
func CheckRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil role dari header (misalnya Authorization atau custom header)
		userRole := c.GetHeader("Role") // Header ini harus dikirim oleh klien

		// Periksa apakah role user sesuai dengan parameter
		if strings.ToLower(userRole) != strings.ToLower(role) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have the required role"})
			c.Abort() // Hentikan eksekusi jika role tidak sesuai
			return
		}

		
		// Lanjutkan jika role sesuai
		c.Next()
	}
}

// Middleware to check user type (jenis_user)
func CheckJenisUser(ju string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil jenis_user dari header (misalnya Authorization atau custom header)
		userJenis := c.GetHeader("Jenis-User") // Header ini harus dikirim oleh klien

		// Periksa apakah jenis_user sesuai dengan parameter
		if strings.ToLower(userJenis) != strings.ToLower(ju) {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have the required user type"})
			c.Abort() // Hentikan eksekusi jika jenis user tidak sesuai
			return
		}

		// Lanjutkan jika jenis_user sesuai
		c.Next()
	}
}


