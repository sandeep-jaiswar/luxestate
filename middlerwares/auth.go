package middleware

import (
	"fmt"
	"luxestate/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Hello, World!")
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			utils.ErrorResponse(c, "Authorization header is required", 401)
			c.Abort()
			return
		}

		// Verify JWT token here (similar to the previous code's authMiddleware)
		// ...

		// Assuming token validation logic here, for example:
		// if valid {
		//     c.Next()
		// } else {
		//     utils.ErrorResponse(c, "Invalid token", 401)
		//     c.Abort()
		//     return
		// }
	}
}
