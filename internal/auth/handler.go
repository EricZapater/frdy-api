package auth

import (
	"frdy-api/internal/users"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
    authService    AuthService
    jwtMiddleware *jwt.GinJWTMiddleware
}

func NewAuthHandler(authService AuthService, jwtMiddleware *jwt.GinJWTMiddleware) *AuthHandler {
    return &AuthHandler{
        authService:    authService,
        jwtMiddleware: jwtMiddleware,
    }
}

// Login godoc
// @Summary User login
// @Description Authenticates a user and returns a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} LoginResponse "Login successful"
// @Failure 400 {object} map[string]string "Invalid login request"
// @Failure 401 {object} map[string]string "Invalid credentials"
// @Failure 403 {object} map[string]string "User inactive"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
    var loginRequest LoginRequest
    if err := c.ShouldBindJSON(&loginRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login request"})
        return
    }
    
    token, user, expire, err := h.authService.Login(c.Request.Context(), loginRequest)
    if err != nil {
        var statusCode int
        switch err {
        case ErrInvalidCredentials:
            statusCode = http.StatusUnauthorized
        case users.ErrInactiveUser:
            statusCode = http.StatusForbidden
        default:
            statusCode = http.StatusInternalServerError
        }
        
        c.JSON(statusCode, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, LoginResponse{
        Token:  token,
        Expire: expire.Format(time.RFC3339),
        User:   user,
    })
}

// RefreshToken godoc
// @Summary Refresh JWT token
// @Description Refreshes an existing JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} map[string]interface{} "Token refreshed successfully"
// @Failure 401 {object} map[string]string "Invalid or expired token"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /auth/refresh_token [get]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
    h.jwtMiddleware.RefreshHandler(c)
}
