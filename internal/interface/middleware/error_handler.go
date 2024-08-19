package middleware

import (
    "restapi/pkg/errors"
    "github.com/gin-gonic/gin"
    "net/http"
)

func ErrorHandler() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        ctx.Next()

        // Check if any errors were set
        if len(ctx.Errors) > 0 {
            err := ctx.Errors.Last().Err

            // Check for custom HTTP errors
            if httpErr, ok := err.(*errors.HTTPError); ok {
                ctx.JSON(httpErr.StatusCode, gin.H{"error": httpErr.Message})
                return
            }

            // Fallback to internal server error for any unhandled errors
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
        }
    }
}
