package controllers

import (
    "admin-dashboard/initializers"
    "admin-dashboard/models"
    "net/http"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
)

// GetAllAnnouncements retrieves all announcements (optionally with pagination)
func GetAllAnnouncements(c *gin.Context) {
    var announcements []models.Announcement
    // Pagination parameters (optional)
    limit := 10
    page := 1

    if l := c.Query("limit"); l != "" {
        limit, _ = strconv.Atoi(l)
    }
    if p := c.Query("page"); p != "" {
        page, _ = strconv.Atoi(p)
    }

    offset := (page - 1) * limit

    if err := initializers.DB.Limit(limit).Offset(offset).Find(&announcements).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, announcements)
}

// PostAnnouncement creates a new announcement
func PostAnnouncement(c *gin.Context) {
    var input struct {
        Title     string    `json:"title" binding:"required"`
        Body      string    `json:"body" binding:"required"`
        StartDate time.Time `json:"start_date" binding:"required"`
        EndDate   time.Time `json:"end_date" binding:"required"`
        IsActive  bool      `json:"is_active" binding:"required"`
    }

    // Validate input
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create the announcement
    announcement := models.Announcement{
        Title:     input.Title,
        Body:      input.Body,
        StartDate: input.StartDate,
        EndDate:   input.EndDate,
        IsActive:  input.IsActive,
    }

    // Save to database
    if err := initializers.DB.Create(&announcement).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, announcement)
}

// DeleteAnnouncement deletes an announcement by ID
func DeleteAnnouncement(c *gin.Context) {
    // Get ID from the URL
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid announcement ID"})
        return
    }

    // Check if the announcement exists
    var announcement models.Announcement
    if err := initializers.DB.First(&announcement, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Announcement not found"})
        return
    }

    // Delete the announcement
    if err := initializers.DB.Delete(&announcement).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Announcement deleted successfully"})
}

