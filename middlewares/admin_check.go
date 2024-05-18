// middlewares/admin_check.go
package middlewares

import (
	"open-contribute/services"

	"github.com/gin-gonic/gin"
)

func AdminCheckMiddleware(service services.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// userID := ctx.MustGet("userID").(uint)
		// user, err := service.GetUserByID(userID)
		// if err != nil || !user.IsAdmin {
		// 	ctx.JSON(http.StatusForbidden, gin.H{"error": "admin access required"})
		// 	ctx.Abort()
		// 	return
		// }
		ctx.Next()
	}
}
