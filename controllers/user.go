package controllers

import (
	"database/sql"
	"encoding/hex"
	"luxestate/db"
	"luxestate/models"
	"luxestate/utils"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Get all users
func GetUsers(c *gin.Context) {
	page := c.Query("page") // Retrieve the page number from the query parameter (if provided)
	limit := 10              // Set a default limit for the number of users per page
	offset := 0              // Initialize offset for pagination

	// Convert the page number to an integer
	pageNumber, err := strconv.Atoi(page)
	if err != nil || pageNumber < 1 {
		pageNumber = 1 // Set default page number if invalid or not provided
	}

	if pageNumber > 1 {
		offset = (pageNumber - 1) * limit // Calculate offset based on the page number
	}

	var users []models.User
	dbInstance := db.GetDBInstance()

	// Query the database to fetch users with pagination
    query := "SELECT id, email, password, salt FROM users LIMIT ? OFFSET ?"
	rows, err := dbInstance.Query(query, limit, offset)
	if err != nil {
		utils.ErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Iterate through the rows and populate the users slice
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.UID, &user.Email, &user.Password, &user.Salt)
		if err != nil {
			utils.ErrorResponse(c, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		utils.ErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the users as a JSON response
	c.JSON(http.StatusOK, users)
}


func GetUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	dbInstance := db.GetDBInstance()

	// Query the database to fetch the user by ID
	err := dbInstance.QueryRow("SELECT id, email, password, salt FROM users WHERE id = ?", id).Scan(&user.UID, &user.Email, &user.Password, &user.Salt)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.ErrorResponse(c, "User not found", http.StatusNotFound)
			return
		}
		utils.ErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the user details as a JSON response
	c.JSON(http.StatusOK, user)
}


func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	salt, err := utils.GenerateSalt(16)
    if err != nil {
        utils.ErrorResponse(c, "Error generating salt", http.StatusInternalServerError)
        return
    }

	saltedPassword := append([]byte(user.Password), salt...)

	hashedPassword, err := bcrypt.GenerateFromPassword(saltedPassword, bcrypt.DefaultCost)
    if err != nil {
        utils.ErrorResponse(c, err.Error(), http.StatusInternalServerError)
        return
    }

	dbInstance := db.GetDBInstance()

	_, err = dbInstance.Exec("INSERT INTO users (email, password, salt) VALUES (?, ?, ?)", user.Email, string(hashedPassword), hex.EncodeToString(salt))
    if err != nil {
        utils.ErrorResponse(c, err.Error(), http.StatusInternalServerError)
        return
    }

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
