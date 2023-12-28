package controllers

import (
	"encoding/hex"
	"luxestate/db"
	"luxestate/utils"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
    var loginRequest struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.BindJSON(&loginRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    dbInstance := db.GetDBInstance()
    var storedPassword, storedSalt string

    // Fetch user's stored hashed password and salt from the database
    err := dbInstance.QueryRow("SELECT password, salt FROM users WHERE email = ?", loginRequest.Email).Scan(&storedPassword, &storedSalt)
    if err != nil {
        utils.ErrorResponse(c, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    // Combine the provided password with the stored salt for comparison
    salt, _ := hex.DecodeString(storedSalt)
    saltedPassword := append([]byte(loginRequest.Password), salt...)

    // Compare hashed password with stored hashed password
    if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), saltedPassword); err != nil {
        utils.ErrorResponse(c, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    // Generate JWT token upon successful authentication
    token, err := utils.GenerateJWTToken(loginRequest.Email) // Implement a function to generate JWT token
    if err != nil {
        utils.ErrorResponse(c, "Error generating token", http.StatusInternalServerError)
        return
    }

    // Return the JWT token in the response
    c.JSON(http.StatusOK, gin.H{"token": token})
}

