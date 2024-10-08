package api

import (
	"admin-dashboard/controllers"
	"admin-dashboard/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	// Public routes
	r.GET("/users", controllers.GetAllUsers)
	r.POST("/users/signup", controllers.Signup)
	r.POST("/users/login", controllers.Login)

	r.GET("/users/me", controllers.GetCurrentUser)


	// User routes (protected)
	r.GET("/users/validate",middleware.RequireAuth, controllers.Validate) // Validate user
	r.PUT("/users/:id", middleware.RequireAuth,controllers.UpdateUser)     // Update user
	r.DELETE("/users/:id",middleware.RequireAuth, controllers.DeleteUser)  // Delete user

	//Announcement Routes
	r.GET("/announcements", controllers.GetAllAnnouncements)
	r.POST("/announcements",controllers.PostAnnouncement)
	r.POST("/announcements/:id",controllers.DeleteAnnouncement)

	//Transaction Routes
	r.POST("/transactions", controllers.CreateTransaction)
	r.GET("/transactions/:id", controllers.GetTransaction)
	r.GET("/transactions", controllers.GetTransactions)
	r.PUT("/transactions/:id", controllers.UpdateTransaction)
	r.DELETE("/transactions/:id", controllers.DeleteTransaction)
	r.GET("/transactions/user/:user", controllers.GetTransactionsByUser)
}

