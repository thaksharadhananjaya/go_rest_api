package controller

import (
	"net/http"
	"restapi/internal/domain/user"
	"restapi/internal/usecase"
	"restapi/pkg/errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
    userUseCase *usecase.UserUseCase
}

// NewUserController creates a new instance of UserController
func NewUserController(uc *usecase.UserUseCase) *UserController {
    return &UserController{
        userUseCase: uc,
    }
}

// CreateUser handles POST requests to create a new user
func (c *UserController) CreateUser(ctx *gin.Context) {
    var userRequest user.User
    if err := ctx.ShouldBindJSON(&userRequest); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    createdUser, err := c.userUseCase.CreateUser(userRequest)
    if err != nil {
        switch err {
        case errors.ErrNotFound:
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        case errors.ErrInternalServerError:
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        default:
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
        }
        return
    }
    ctx.JSON(http.StatusOK, gin.H{
        "message": "User created successfully",
        "user":    createdUser,
    })
}

// GetUserByID handles GET requests to retrieve a user by ID
func (c *UserController) GetUserByID(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    user, err := c.userUseCase.GetUserByID(uint(id))
    if err != nil {
        switch err {
        case errors.ErrNotFound:
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        case errors.ErrInternalServerError:
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        default:
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
        }
        return
    }

    ctx.JSON(http.StatusOK, user)
}

// GetAllUsers handles GET requests to retrieve all users
func (c *UserController) GetAllUsers(ctx *gin.Context) {
    users, err := c.userUseCase.GetAllUsers()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, users)
}

// UpdateUser handles PUT requests to update an existing user
func (c *UserController) UpdateUser(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    var updateData map[string]interface{}
    if err := ctx.ShouldBindJSON(&updateData); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    updatedUser, err := c.userUseCase.UpdateUser(uint(id), updateData)
    if err != nil {
        switch err {
        case errors.ErrNotFound:
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        case errors.ErrInternalServerError:
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        default:
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
        }
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": updatedUser})
}

// DeleteUser handles DELETE requests to delete a user by ID
func (c *UserController) DeleteUser(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    err = c.userUseCase.DeleteUser(uint(id))
    if err != nil {
        switch err {
        case errors.ErrNotFound:
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        case errors.ErrInternalServerError:
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        default:
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
        }
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}


