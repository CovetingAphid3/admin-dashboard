package controllers

import (
	"admin-dashboard/initializers"
	"admin-dashboard/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(c *gin.Context) {
    var users []models.User

    // Fetch all users from the database
    if err := initializers.DB.Find(&users).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Return the users as a JSON response
    c.JSON(http.StatusOK, users)
}

func Signup(c *gin.Context) {
	// get email, password, first name, and last name from request body
	var body struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// check if email or password is empty
	if body.Email == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or password cannot be empty"})
		return
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// create the user with a default role and additional fields
	user := models.User{
		Email:     body.Email,
		Password:  string(hash),
		Role:      "user",
		FirstName: body.FirstName,
		LastName:  body.LastName,
		IsActive:  true, // set IsActive to true by default
		LastLogin: time.Time{}, // set to zero value (no last login)
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}

	// respond with success
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}



func Login(c *gin.Context) {
	// get email and password from request body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// check if email or password is empty
	if body.Email == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or password cannot be empty"})
		return
	}

	// look up the user by email
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// compare the provided password with the stored hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// set cookie with the token
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	// respond with success
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// UpdateUser updates a user's information
func UpdateUser(c *gin.Context) {
    var user models.User
    userID := c.Param("id") // Get user ID from URL parameter

    // Check if user exists
    if err := initializers.DB.First(&user, userID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Bind request body
    var body struct {
        Email    string
        Password string
    }

    if err := c.Bind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    // Update user fields
    if body.Email != "" {
        user.Email = body.Email
    }
    if body.Password != "" {
        hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
            return
        }
        user.Password = string(hash)
    }

    // Save updated user
    if err := initializers.DB.Save(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser deletes a user
func DeleteUser(c *gin.Context) {
    userID := c.Param("id") // Get user ID from URL parameter

    // Check if user exists
    if err := initializers.DB.Delete(&models.User{}, userID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}


func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{"message": user})
}
